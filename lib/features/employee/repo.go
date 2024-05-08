package employee

type EmployeesRepository interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
}

func NewEmployeeRepo() EmployeesRepository {
	return &employeeRepo{}
}

type employeeRepo struct {
}

func (r *employeeRepo) CreateEmployee(input *CreateEmployeeInputDto) (*EmployeeDto, error) {
	return &EmployeeDto{Password: input.Password}, nil
}
