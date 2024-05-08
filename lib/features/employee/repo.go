package employee




type EmployeesRepository interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
}
