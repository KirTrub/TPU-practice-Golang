package models

import "time"

type Hospitalization struct {
	ID                  int32     `json:"id"`
	PatientID           int32     `json:"patient_id"`
	DoctorID            int32     `json:"doctor_id"`
	DiagnosisID         int32     `json:"diagnosis_id"`
	DepartamentID       int32     `json:"departament_id"`
	StartDate           string    `json:"start_date"`  // Формат 'YYYY-MM-DD'
	FinishDate          string    `json:"finish_date"` // Формат 'YYYY-MM-DD'
	HospitalizationDate time.Time `json:"hospitalization_date,omitempty"`
}

type HospitalizationResponse struct {
	ID                  int32     `json:"id"`
	StartDate           string    `json:"start_date"`
	FinishDate          string    `json:"finish_date"`
	HospitalizationDate time.Time `json:"hospitalization_date"`

	Patient     PatientInfo     `json:"patient"`
	Doctor      DoctorInfo      `json:"doctor"`
	Diagnosis   DiagnosisInfo   `json:"diagnosis"`
	Departament DepartamentInfo `json:"departament"`
}

type PatientInfo struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DoctorInfo struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DiagnosisInfo struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type DepartamentInfo struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}
