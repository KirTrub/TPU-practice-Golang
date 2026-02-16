package models

type Doctor struct {
	ID            int32   `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	SurName       *string `json:"sur_name,omitempty"`
	DepartamentID int32   `json:"departament_id"`
}

type DoctorResponse struct {
	ID               int32   `json:"id"`
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	SurName          *string `json:"sur_name,omitempty"`
	DepartamentID    int32   `json:"departament_id"`
	DepartamentTitle string  `json:"departament_title"`
}
