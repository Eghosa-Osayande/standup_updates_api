package sprint

import (
	"standup-api/lib/common/dto"
	"time"
)

type CreateSprintInputDto struct {
	Name             string    `json:"name" validate:"required"`
	StandupStartTime *dto.Json24HrTime `json:"standup_start_time" validate:"required"`
	StandupEndTime   *dto.Json24HrTime `json:"standup_end_time" validate:"required"`
}



type SprintDto struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"created_at"`
	StandupStartTime *dto.Json24HrTime `json:"standup_start_time"`
	StandupEndTime   *dto.Json24HrTime `json:"standup_end_time"`
}

type FetchSprintsInputDto struct {
	Page int `json:"page" validate:"required"`
	PerPage int `json:"per_page" validate:"required"`
}
