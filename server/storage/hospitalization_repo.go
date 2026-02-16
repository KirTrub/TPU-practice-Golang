package storage

import (
	"context"
	"hospital-app/server/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HospitalizationRepository struct {
	db *pgxpool.Pool
}

func NewHospitalizationRepository(db *pgxpool.Pool) *HospitalizationRepository {
	return &HospitalizationRepository{db: db}
}

func (r *HospitalizationRepository) CreateHospitalization(ctx context.Context, h *models.Hospitalization) (int32, error) {
	var id int32
	query := `
		INSERT INTO hospitalization
			(number_patient, id_doctor, id_diagnosis, id_departament,
			 start_hospitalization, finish_hospitalization, hospitalization_date)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id_hospitalization`

	err := r.db.QueryRow(ctx, query,
		h.PatientID, h.DoctorID, h.DiagnosisID, h.DepartamentID, h.StartDate, h.FinishDate,
	).Scan(&id)

	return id, err
}

func (r *HospitalizationRepository) GetAllHospitalizations(ctx context.Context) ([]models.HospitalizationResponse, error) {
	var hospitalizations []models.HospitalizationResponse
	query := `
		SELECT
			h.id_hospitalization, h.start_hospitalization, h.finish_hospitalization, h.hospitalization_date,
			p.number_patient, p.first_name, p.second_name,
			d.id_doctor, d.first_name, d.second_name,
			diag.id_diagnosis, diag.title_diagnosis,
			dep.id_departament, dep.title_departament
		FROM hospitalization h
		JOIN patient p ON h.number_patient = p.number_patient
		JOIN doctor d ON h.id_doctor = d.id_doctor
		JOIN diagnosis diag ON h.id_diagnosis = diag.id_diagnosis
		JOIN departament dep ON h.id_departament = dep.id_departament
		ORDER BY h.hospitalization_date DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h models.HospitalizationResponse
		var startDate, finishDate time.Time
		if err := rows.Scan(
			&h.ID, &startDate, &finishDate, &h.HospitalizationDate,
			&h.Patient.ID, &h.Patient.FirstName, &h.Patient.LastName,
			&h.Doctor.ID, &h.Doctor.FirstName, &h.Doctor.LastName,
			&h.Diagnosis.ID, &h.Diagnosis.Title,
			&h.Departament.ID, &h.Departament.Title,
		); err != nil {
			return nil, err
		}
		h.StartDate = startDate.Format("2006-01-02")
		h.FinishDate = finishDate.Format("2006-01-02")
		hospitalizations = append(hospitalizations, h)
	}

	return hospitalizations, nil
}

func (r *HospitalizationRepository) DeleteHospitalization(ctx context.Context, id int32) error {
	query := `DELETE FROM hospitalization WHERE id_hospitalization = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
