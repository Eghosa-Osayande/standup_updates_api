package employee

import (
	"log"
	"net/http"
	"os"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/jwt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, *http_response.HttpError)

	Login(*EmployeeLoginInputDto) (*http_response.HttpResponse[string], *http_response.HttpError)

	FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.HttpPagedResponse[EmployeeDto], *http_response.HttpError)
}

func NewEmployeeService(repo EmployeesRepository) EmployeeService {
	return &employeeService{employeeRepo: repo}
}

type employeeService struct {
	employeeRepo EmployeesRepository
}

func (s *employeeService) CreateEmployee(input *CreateEmployeeInputDto) (*EmployeeDto, *http_response.HttpError) {

	emp, _ := s.employeeRepo.GetEmployeeByName(input.Name)
	if emp != nil {
		return nil, http_response.NewHttpError(http.StatusBadRequest, "employee already exists")
	}

	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		log.Println(hashErr)
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error creating employee")
	}

	input.Password = string(hashedPass)

	employee, err := s.employeeRepo.CreateEmployee(input)
	if err != nil {
		log.Println(err)
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error creating employee")
	}

	return employee, nil
}

func (s *employeeService) Login(input *EmployeeLoginInputDto) (*http_response.HttpResponse[string], *http_response.HttpError) {

	employee, _ := s.employeeRepo.GetEmployeeByName(input.Name)
	if employee==nil {
		return nil, http_response.NewHttpError(http.StatusUnauthorized, "employee does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(input.Password))
	

	if err != nil {
		return nil, http_response.NewHttpError(http.StatusUnauthorized, "invalid credentials")
	}
	secret, ok := os.LookupEnv("TOKEN_KEY")
	if !ok {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error generating token")
	}

	token, tokenErr := jwt.GenerateJwtToken(map[string]any{"role": "employee","id":employee.ID}, secret, 3*time.Hour)

	if tokenErr != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error generating token")
	}
	response := http_response.NewSuccessResponse(token)
	return &response, nil
}

func (s *employeeService) FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.HttpPagedResponse[EmployeeDto], *http_response.HttpError) {
	employees, err := s.employeeRepo.FindAllEmployees(input)

	if err != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error fetching employees")
	}



	result := http_response.NewPagedResponse(*employees)

	return &result, nil
}
