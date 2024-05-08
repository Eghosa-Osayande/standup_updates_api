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
    employees ( name, password)
values
    ( $1, $2) Returning count_id, id, created_at, updated_at, name, password
`

type CreateEmployeeParams struct {
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRow(ctx, createEmployee, arg.Name, arg.Password)
	var i Employee
	err := row.Scan(
		&i.CountID,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Password,
	)
	return i, err
}

const fetchAllEmployees = `-- name: FetchAllEmployees :many
Select
    count_id, id, created_at, updated_at, name, password
from
    employees

ORDER BY
    employees.updated_at DESC,
    employees.count_id DESC

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
			&i.CountID,
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Password,
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
    count_id, id, created_at, updated_at, name, password
from
    employees
where
    name = $1
`

func (q *Queries) FetchEmployeeByName(ctx context.Context, name string) (Employee, error) {
	row := q.db.QueryRow(ctx, fetchEmployeeByName, name)
	var i Employee
	err := row.Scan(
		&i.CountID,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Password,
	)
	return i, err
}
