package models

type HospitalizationReport struct {
    Diagnosis     string `json:"diagnosis"`
    DoctorFIO     string `json:"doctor_fio"`
    PatientCount  int32  `json:"patient_count"`
    MinDays       int32  `json:"min_days"`
    MaxDays       int32  `json:"max_days"`
}
