-- name: CreateEmployee :one
Insert into
    employees ( name, password)
values
    ( $1, $2) Returning *;

-- name: FetchAllEmployees :many
Select
    *
from
    employees

ORDER BY
    employees.created_at DESC

limit
    sqlc.narg('limit')
offset
    sqlc.narg('offset');

-- name: FetchEmployeeByName :one
Select
    *
from
    employees
where
    name = $1;