-- name: CreateSprint :one
Insert into
    sprints (name, standup_start_time, standup_end_time)
values
    ( $1, $2, $3) Returning *;

-- name: FetchAllSprint :many
Select
    *
from
    sprints

ORDER BY
    sprints.created_at DESC
limit
    sqlc.narg('limit')
offset
    sqlc.narg('offset');

-- name: FetchSprintById :one
Select
    *
from
    sprints
where
    id = $1;