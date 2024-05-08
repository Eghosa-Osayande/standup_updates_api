// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sprint.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSprint = `-- name: CreateSprint :one
Insert into
    sprints ( name, standup_start_time, standup_end_time)
values
    ( $1, $2, $3) Returning id, created_at, name, standup_start_time, standup_end_time
`

type CreateSprintParams struct {
	Name             string             `db:"name" json:"name"`
	StandupStartTime pgtype.Timestamptz `db:"standup_start_time" json:"standup_start_time"`
	StandupEndTime   pgtype.Timestamptz `db:"standup_end_time" json:"standup_end_time"`
}

func (q *Queries) CreateSprint(ctx context.Context, arg CreateSprintParams) (Sprint, error) {
	row := q.db.QueryRow(ctx, createSprint, arg.Name, arg.StandupStartTime, arg.StandupEndTime)
	var i Sprint
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.StandupStartTime,
		&i.StandupEndTime,
	)
	return i, err
}

const fetchAllSprint = `-- name: FetchAllSprint :many
Select
    id, created_at, name, standup_start_time, standup_end_time
from
    sprints

ORDER BY
    sprints.created_at DESC
limit
    $2
offset
    $1
`

type FetchAllSprintParams struct {
	Offset pgtype.Int4 `db:"offset" json:"offset"`
	Limit  pgtype.Int4 `db:"limit" json:"limit"`
}

func (q *Queries) FetchAllSprint(ctx context.Context, arg FetchAllSprintParams) ([]Sprint, error) {
	rows, err := q.db.Query(ctx, fetchAllSprint, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sprint
	for rows.Next() {
		var i Sprint
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.StandupStartTime,
			&i.StandupEndTime,
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

const fetchSprintById = `-- name: FetchSprintById :one
Select
    id, created_at, name, standup_start_time, standup_end_time
from
    sprints
where
    id = $1
`

func (q *Queries) FetchSprintById(ctx context.Context, id pgtype.UUID) (Sprint, error) {
	row := q.db.QueryRow(ctx, fetchSprintById, id)
	var i Sprint
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.StandupStartTime,
		&i.StandupEndTime,
	)
	return i, err
}