package storage

import (
	"context"
	"hospital-app/server/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepository struct {
	db *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetHospitalizationReport(
	ctx context.Context,
	departmentID int,
	year int,
) ([]models.HospitalizationReport, error) {

	query := `
        SELECT
    diag.title_diagnosis AS diagnosis,
    CONCAT(
        doc.second_name, ' ',
        doc.first_name, ' ',
        COALESCE(doc.sur_name, '')
    ) AS doctor_fio,
    COUNT(DISTINCT h.number_patient) AS patient_count,
    MIN(h.finish_hospitalization - h.start_hospitalization) AS min_days,
    MAX(h.finish_hospitalization - h.start_hospitalization) AS max_days
FROM hospitalization h
JOIN diagnosis diag ON h.id_diagnosis = diag.id_diagnosis
JOIN doctor doc ON h.id_doctor = doc.id_doctor
WHERE h.id_departament = $1
  AND EXTRACT(YEAR FROM h.hospitalization_date) = $2
GROUP BY
    diag.title_diagnosis,
    doc.second_name,
    doc.first_name,
    doc.sur_name
ORDER BY
    diag.title_diagnosis,
    doctor_fio;

    `

	rows, err := r.db.Query(ctx, query, departmentID, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.HospitalizationReport
	for rows.Next() {
		var r models.HospitalizationReport
		if err := rows.Scan(
			&r.Diagnosis,
			&r.DoctorFIO,
			&r.PatientCount,
			&r.MinDays,
			&r.MaxDays,
		); err != nil {
			return nil, err
		}
		res = append(res, r)
	}

	return res, nil
}
