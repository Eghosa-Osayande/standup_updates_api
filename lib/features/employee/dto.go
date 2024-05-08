package employee

import "time"

type CreateEmployeeInputDto struct {
	Name string `json:"name"`
}

type EmployeeDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
