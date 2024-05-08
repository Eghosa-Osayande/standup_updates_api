-- name: CreateUpdate :one
Insert into
    updates (
        employee_id,
        sprint_id,
        yesterday,
        today,
        blocked_by,
        breakaway,
        check_in_time,
        status,
        tasks
    )
values
    ($1, $2, $3, $4, $5, $6, $7, $8, $9) Returning *;

-- name: FetchAllUpdates :many
Select
    updates.*,
    employees.name as employee_name
from
    updates
    join employees on updates.employee_id = employees.id
    join sprints on updates.sprint_id = sprints.id
where
    (
        updates.sprint_id = sqlc.narg('sprint_id')
        or sqlc.narg('sprint_id') IS NULL
    )
    AND
    (
        updates.employee_id = sqlc.narg('employee_id')
        or sqlc.narg('employee_id') IS NULL
    )
    AND
    (
        updates.created_at >= sqlc.narg('start_date')
        or sqlc.narg('start_date') IS NULL
    )
    AND
    (
        updates.created_at <= sqlc.narg('end_date')
        or sqlc.narg('end_date') IS NULL
    )
ORDER BY
    sprints.created_at DESC
limit
    sqlc.narg('limit') offset sqlc.narg('offset');