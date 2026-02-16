package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(
	patientHandler *PatientHandler,
	departamentHandler *DepartamentHandler,
	diagnosisHandler *DiagnosisHandler,
	doctorHandler *DoctorHandler,
	hospitalizationHandler *HospitalizationHandler,
	reportHandler *ReportHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hospital API is running!"))
	}).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	// Пациенты
	api.HandleFunc("/patients", patientHandler.CreatePatientHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/patients", patientHandler.GetAllPatientsHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/patients/{id:[0-9]+}", patientHandler.GetPatientByIDHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/patients/{id:[0-9]+}", patientHandler.UpdatePatientHandler).Methods("PUT", "OPTIONS")
	api.HandleFunc("/patients/{id:[0-9]+}", patientHandler.DeletePatientHandler).Methods("DELETE", "OPTIONS")

	// Отделения
	api.HandleFunc("/departments", departamentHandler.CreateDepartamentHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/departments", departamentHandler.GetAlldepartmentsHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/departments/{id:[0-9]+}", departamentHandler.UpdateDepartamentHandler).Methods("PUT", "OPTIONS")
	api.HandleFunc("/departments/{id:[0-9]+}", departamentHandler.DeleteDepartamentHandler).Methods("DELETE", "OPTIONS")

	// Диагнозы
	api.HandleFunc("/diagnoses", diagnosisHandler.CreateDiagnosisHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/diagnoses", diagnosisHandler.GetAllDiagnosesHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/diagnoses/{id:[0-9]+}", diagnosisHandler.UpdateDiagnoseHandler).Methods("PUT", "OPTIONS")
	api.HandleFunc("/diagnoses/{id:[0-9]+}", diagnosisHandler.DeleteDiagnoseHandler).Methods("DELETE", "OPTIONS")

	// Врачи
	api.HandleFunc("/doctors", doctorHandler.CreateDoctorHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/doctors", doctorHandler.GetAllDoctorsHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/doctors/{id:[0-9]+}", doctorHandler.UpdateDoctorHandler).Methods("PUT", "OPTIONS")
	api.HandleFunc("/doctors/{id:[0-9]+}", doctorHandler.DeleteDoctorHandler).Methods("DELETE", "OPTIONS")

	// Госпитализации
	api.HandleFunc("/hospitalizations", hospitalizationHandler.CreateHospitalizationHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/hospitalizations", hospitalizationHandler.GetAllHospitalizationsHandler).Methods("GET", "OPTIONS")
	api.HandleFunc("/hospitalizations/{id:[0-9]+}", hospitalizationHandler.DeleteHospitalizationHandler).Methods("DELETE", "OPTIONS")

	//Отчет
	api.HandleFunc("/reports/hospitalizations", reportHandler.GetHospitalizationReport).Methods("GET", "OPTIONS")

	return r
}
