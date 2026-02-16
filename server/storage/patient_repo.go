package storage

import (
	"context"
	"hospital-app/server/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PatientRepository struct {
	db *pgxpool.Pool
}

func NewPatientRepository(db *pgxpool.Pool) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) CreatePatient(ctx context.Context, patient *models.Patient) (int32, error) {
	var id int32
	query := `INSERT INTO patient (first_name, second_name, sur_name, gender, date_of_birth, patient_address)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING number_patient`
	err := r.db.QueryRow(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.SurName,
		patient.Gender,
		patient.BirthDate,
		patient.Address,
	).Scan(&id)
	return id, err
}

func (r *PatientRepository) GetAllPatients(ctx context.Context) ([]models.Patient, error) {
	var patients []models.Patient
	query := `SELECT number_patient, first_name, second_name, sur_name, gender, date_of_birth, patient_address FROM patient`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Patient
		var birthDate time.Time
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.SurName, &p.Gender, &birthDate, &p.Address); err != nil {
			return nil, err
		}
		p.BirthDate = birthDate.Format("2006-01-02")
		patients = append(patients, p)
	}
	return patients, nil
}

func (r *PatientRepository) GetPatientByID(ctx context.Context, id int32) (*models.Patient, error) {
	var p models.Patient
	var birthDate time.Time
	query := `SELECT number_patient, first_name, second_name, sur_name, gender, date_of_birth, patient_address
			  FROM patient WHERE number_patient = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.SurName, &p.Gender, &birthDate, &p.Address)
	if err != nil {
		return nil, err
	}
	p.BirthDate = birthDate.Format("2006-01-02")
	return &p, nil
}

func (r *PatientRepository) UpdatePatient(ctx context.Context, patient *models.Patient) error {
	query := `UPDATE patient SET first_name = $1, second_name = $2, sur_name = $3,
			  gender = $4, date_of_birth = $5, patient_address = $6
			  WHERE number_patient = $7`
	_, err := r.db.Exec(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.SurName,
		patient.Gender,
		patient.BirthDate,
		patient.Address,
		patient.ID,
	)
	return err
}

func (r *PatientRepository) DeletePatient(ctx context.Context, id int32) error {
	query := `DELETE FROM patient WHERE number_patient = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
