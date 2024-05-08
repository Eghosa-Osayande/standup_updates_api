package sprint

import "time"

type CreateSprintInputDto struct {
	Name             string    `json:"name"`
	StartedAt        time.Time `json:"started_at"`
	EndedAt          time.Time `json:"ended_at"`
	StandupStartTime time.Time `json:"standup_start_time"`
	StandupEndTime   time.Time `json:"standup_end_time"`
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
