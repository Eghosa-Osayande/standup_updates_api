package employee

import (
	"context"
	"standup-api/lib/common/database"
	"standup-api/lib/utils/http_response"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmployeesRepository interface {
	CreateEmployee(*CreateEmployeeInputDto) (*EmployeeDto, error)
	GetEmployeeByName(string) (*EmployeeDto, error)
	FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.CursorPage[EmployeeDto], error)
}

func NewEmployeeRepo(db *database.Database) EmployeesRepository {
	return &employeeRepo{
		db: db,
	}
}

type employeeRepo struct {
	db *database.Database
}

func (r *employeeRepo) CreateEmployee(input *CreateEmployeeInputDto) (*EmployeeDto, error) {
	employee, err := r.db.CreateEmployee(context.Background(), database.CreateEmployeeParams{
		Name:     input.Name,
		Password: input.Password,
	})

	if err != nil {
		return nil, err
	}

	return EmployeeModelToDto(&employee), nil
}

func (r *employeeRepo) GetEmployeeByName(name string) (*EmployeeDto, error) {
	employee, err := r.db.FetchEmployeeByName(context.Background(), name)

	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return EmployeeModelToDto(&employee), nil
}

func (r *employeeRepo) FindAllEmployees(input *FetchEmployeesInputDto) (*http_response.CursorPage[EmployeeDto], error) {
	emps, err := r.db.FetchAllEmployees(context.Background(), database.FetchAllEmployeesParams{
		Offset: pgtype.Int4{Int32: int32((input.Page - 1) * input.PerPage), Valid: true},
		Limit:  pgtype.Int4{Int32: int32(input.PerPage), Valid: true},
	})
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	var empsSlice []EmployeeDto = make([]EmployeeDto, 0)

	for _, emp := range emps {
		empsSlice = append(empsSlice, *EmployeeModelToDto(&emp))

	}

	return &http_response.CursorPage[EmployeeDto]{Data: empsSlice, Cursor: &http_response.Cursor{NextPage: input.Page + 1, PerPage: input.PerPage}}, nil
}
