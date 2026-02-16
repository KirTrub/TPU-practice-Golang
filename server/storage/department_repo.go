package storage

import (
	"context"
	"hospital-app/server/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartamentRepository struct {
	db *pgxpool.Pool
}

func NewDepartamentRepository(db *pgxpool.Pool) *DepartamentRepository {
	return &DepartamentRepository{db: db}
}

func (r *DepartamentRepository) CreateDepartament(ctx context.Context, dep *models.Departament) (int32, error) {
	var id int32
	query := `INSERT INTO departament (title_departament) VALUES ($1) RETURNING id_departament`
	err := r.db.QueryRow(ctx, query, dep.Title).Scan(&id)
	return id, err
}

func (r *DepartamentRepository) GetAlldepartments(ctx context.Context) ([]models.Departament, error) {
	var departments []models.Departament
	query := `SELECT id_departament, title_departament FROM departament`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Departament
		if err := rows.Scan(&d.ID, &d.Title); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}
	return departments, nil
}

func (r *DepartamentRepository) UpdateDepartament(ctx context.Context, dep *models.Departament) error {
	query := `UPDATE departament SET title_departament = $1 WHERE id_departament = $2`
	_, err := r.db.Exec(ctx, query, dep.Title, dep.ID)
	return err
}

func (r *DepartamentRepository) DeleteDepartament(ctx context.Context, id int32) error {
	query := `DELETE FROM departament WHERE id_departament = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
