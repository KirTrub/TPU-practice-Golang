package storage

import (
	"context"
	"hospital-app/server/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DiagnosisRepository struct {
	db *pgxpool.Pool
}

func NewDiagnosisRepository(db *pgxpool.Pool) *DiagnosisRepository {
	return &DiagnosisRepository{db: db}
}

func (r *DiagnosisRepository) CreateDiagnosis(ctx context.Context, diag *models.Diagnosis) (int32, error) {
	var id int32
	query := `INSERT INTO diagnosis (title_diagnosis) VALUES ($1) RETURNING id_diagnosis`
	err := r.db.QueryRow(ctx, query, diag.Title).Scan(&id)
	return id, err
}

func (r *DiagnosisRepository) GetAllDiagnoses(ctx context.Context) ([]models.Diagnosis, error) {
	var diagnoses []models.Diagnosis
	query := `SELECT id_diagnosis, title_diagnosis FROM diagnosis`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Diagnosis
		if err := rows.Scan(&d.ID, &d.Title); err != nil {
			return nil, err
		}
		diagnoses = append(diagnoses, d)
	}
	return diagnoses, nil
}

func (r *DiagnosisRepository) UpdateDiagnose(ctx context.Context, dep *models.Diagnosis) error {
	query := `UPDATE diagnosis SET title_diagnosis = $1 WHERE id_diagnosis = $2`
	_, err := r.db.Exec(ctx, query, dep.Title, dep.ID)
	return err
}

func (r *DiagnosisRepository) DeleteDiagnose(ctx context.Context, id int32) error {
	query := `DELETE FROM diagnosis WHERE id_diagnosis = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
