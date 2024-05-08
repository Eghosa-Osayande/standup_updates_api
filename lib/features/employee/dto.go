package employee

import "time"

type CreateEmployeeInputDto struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EmployeeDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
