package models

type Patient struct {
	ID        int32   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	SurName   *string `json:"sur_name,omitempty"`
	Gender    string  `json:"gender"`
	BirthDate string  `json:"birth_date"` // Формат'YYYY-MM-DD'
	Address   string  `json:"address"`
}
