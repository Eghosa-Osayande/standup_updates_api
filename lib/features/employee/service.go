package employee

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
}

func NewEmployeeService(repo EmployeesRepository) EmployeeService {
	return &employeeService{employeeRepo: repo}
}

type employeeService struct {
	employeeRepo EmployeesRepository
}

func (s *employeeService) CreateEmployee(input *CreateEmployeeInputDto) (*EmployeeDto, error) {
	displayErr := errors.New("error creating employee")
	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		// log
		return nil, displayErr
	}

	input.Password = string(hashedPass)

	employee, err := s.employeeRepo.CreateEmployee(input)
	if err != nil {
		return nil, displayErr
	}

	return employee, nil
}
