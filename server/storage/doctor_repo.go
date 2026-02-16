// server/storage/doctor_repo.go

package storage

import (
	"context"
	"hospital-app/server/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DoctorRepository struct {
	db *pgxpool.Pool
}

func NewDoctorRepository(db *pgxpool.Pool) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doc *models.Doctor) (int32, error) {
	var id int32
	query := `INSERT INTO doctor (first_name, second_name, sur_name, id_departament)
              VALUES ($1, $2, $3, $4) RETURNING id_doctor`
	err := r.db.QueryRow(ctx, query,
		doc.FirstName,
		doc.LastName,
		doc.SurName,
		doc.DepartamentID,
	).Scan(&id)
	return id, err
}

func (r *DoctorRepository) GetAllDoctors(ctx context.Context) ([]models.DoctorResponse, error) {
	var doctors []models.DoctorResponse
	query := `
		SELECT d.id_doctor, d.first_name, d.second_name, d.sur_name, d.id_departament, dep.title_departament
		FROM doctor d
		JOIN departament dep ON d.id_departament = dep.id_departament`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.DoctorResponse
		if err := rows.Scan(&d.ID, &d.FirstName, &d.LastName, &d.SurName, &d.DepartamentID, &d.DepartamentTitle); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}
	return doctors, nil
}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, doc *models.Doctor) error {
	query := `UPDATE doctor SET first_name = $1, second_name = $2, sur_name = $3, id_departament = $4
			  WHERE id_doctor = $5`
	_, err := r.db.Exec(ctx, query,
		doc.FirstName,
		doc.LastName,
		doc.SurName,
		doc.DepartamentID,
		doc.ID,
	)
	return err
}

func (r *DoctorRepository) DeleteDoctor(ctx context.Context, id int32) error {
	query := `DELETE FROM doctor WHERE id_doctor = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
