// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: employee.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEmployee = `-- name: CreateEmployee :one
Insert into
    employees ( name, password,role)
values
    ( $1, $2, $3) Returning id, created_at, name, password, role
`

type CreateEmployeeParams struct {
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRow(ctx, createEmployee, arg.Name, arg.Password, arg.Role)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const fetchAllEmployees = `-- name: FetchAllEmployees :many
Select
    id, created_at, name, password, role
from
    employees

ORDER BY
    employees.created_at DESC

limit
    $2
offset
    $1
`

type FetchAllEmployeesParams struct {
	Offset pgtype.Int4 `db:"offset" json:"offset"`
	Limit  pgtype.Int4 `db:"limit" json:"limit"`
}

func (q *Queries) FetchAllEmployees(ctx context.Context, arg FetchAllEmployeesParams) ([]Employee, error) {
	rows, err := q.db.Query(ctx, fetchAllEmployees, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Password,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchEmployeeByName = `-- name: FetchEmployeeByName :one
Select
    id, created_at, name, password, role
from
    employees
where
    name = $1
`

func (q *Queries) FetchEmployeeByName(ctx context.Context, name string) (Employee, error) {
	row := q.db.QueryRow(ctx, fetchEmployeeByName, name)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Password,
		&i.Role,
	)
	return i, err
}
