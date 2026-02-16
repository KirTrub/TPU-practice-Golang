package main

import (
	"context"
	"hospital-app/server/api"
	"hospital-app/server/storage"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	log.Println("Starting hospital backend server...")

	dbpool, err := storage.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbpool.Close()
	log.Println("Successfully connected to the database")

	migrationSQL, err := os.ReadFile("./db/migration/001_init_schema.sql")
	if err != nil {
		log.Fatalf("Could not read migration file: %v", err)
	}
	_, err = dbpool.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		log.Fatalf("Could not run migration: %v", err)
	}
	log.Println("Database migrations checked/completed successfully")

	patientRepo := storage.NewPatientRepository(dbpool)
	departamentRepo := storage.NewDepartamentRepository(dbpool)
	diagnosisRepo := storage.NewDiagnosisRepository(dbpool)
	doctorRepo := storage.NewDoctorRepository(dbpool)
	hospitalizationRepo := storage.NewHospitalizationRepository(dbpool)

	patientHandler := api.NewPatientHandler(patientRepo)
	departamentHandler := api.NewDepartamentHandler(departamentRepo)
	diagnosisHandler := api.NewDiagnosisHandler(diagnosisRepo)
	doctorHandler := api.NewDoctorHandler(doctorRepo)
	hospitalizationHandler := api.NewHospitalizationHandler(hospitalizationRepo)
	reportRepo := storage.NewReportRepository(dbpool)
	reportHandler := api.NewReportHandler(reportRepo)

	router := api.NewRouter(patientHandler, departamentHandler, diagnosisHandler, doctorHandler, hospitalizationHandler, reportHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
