package employee



type EmployeeService interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
}
