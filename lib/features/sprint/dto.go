package sprint

import (
	"standup-api/lib/common/dto"
	"time"
)

type CreateSprintInputDto struct {
	Name             string    `json:"name" validate:"required"`
	StartedAt        *dto.JsonDate `json:"started_at" validate:"omitnil"`
	EndedAt          *dto.JsonDate `json:"ended_at" validate:"omitnil"`
	StandupStartTime *dto.JsonDate `json:"standup_start_time" validate:"required"`
	StandupEndTime   *dto.JsonDate `json:"standup_end_time" validate:"required"`
}

type UpdateSprintInputDto struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	StartedAt        time.Time `json:"started_at"`
	EndedAt          time.Time `json:"ended_at"`
	StandupStartTime time.Time `json:"standup_start_time"`
	StandupEndTime   time.Time `json:"standup_end_time"`
}

type SprintDto struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	StartedAt        time.Time `json:"started_at"`
	EndedAt          time.Time `json:"ended_at"`
	StandupStartTime time.Time `json:"standup_start_time"`
	StandupEndTime   time.Time `json:"standup_end_time"`
}
type FetchSprintsInputDto struct {
	Cursor string `json:"cursor" validate:"required"`
}
