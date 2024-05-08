
-- name: CreateUpdate :one
Insert into
    updates ( employee_id, sprint_id, yesterday, today, blocked_by, breakaway, check_in_time, status, tasks)
values
    ( $1, $2, $3, $4, $5, $6, $7, $8, $9) Returning *;

-- -- name: FetchAllSprint :many
-- Select
--     *
-- from
--     sprints

-- ORDER BY
--     sprints.created_at DESC
-- limit
--     sqlc.narg('limit')
-- offset
--     sqlc.narg('offset');

