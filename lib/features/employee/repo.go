package employee

import "standup-api/lib/utils/http_response"

type EmployeesRepository interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
	GetEmployeeByName(string) (*EmployeeDto, error)
	FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.CursorPage[EmployeeDto], error)
}

func NewEmployeeRepo() EmployeesRepository {
	return &employeeRepo{}
}

type employeeRepo struct {
}

func (r *employeeRepo) CreateEmployee(input *CreateEmployeeInputDto) (*EmployeeDto, error) {
	return &EmployeeDto{Password: input.Password}, nil
}

func (r *employeeRepo) GetEmployeeByName(name string) (*EmployeeDto, error) {
	return &EmployeeDto{Name: name, Password: ""}, nil
}

func (r *employeeRepo) FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.CursorPage[EmployeeDto], error) {
	return &http_response.CursorPage[EmployeeDto]{}, nil
}
