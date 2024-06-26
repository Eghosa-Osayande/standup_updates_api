package employee

import "time"

type CreateEmployeeInputDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role    string `json:"role" validate:"required,oneof=employee admin"`
}

type EmployeeDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Role    string `json:"role" validate:"required"`
}

type EmployeeLoginInputDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FetchEmployeesInputDto struct {
	Page int `json:"page" validate:"required"`
	PerPage int `json:"per_page" validate:"required"`
}
