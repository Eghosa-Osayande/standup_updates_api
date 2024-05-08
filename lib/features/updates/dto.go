package updates

import (
	"standup-api/lib/common/dto"
	"time"
)

type CreateUpdateInputDto struct {
	EmployeeID  string       `json:"employee_id" validate:"omitempty"`
	SprintID    string       `json:"sprint_id" validate:"required"`
	Yesterday   string       `json:"yesterday" validate:"required"`
	Today       string       `json:"today" validate:"required"`
	BlockedBy   []string     `json:"blocked_by" validate:"required"`
	Tasks   []string     `json:"tasks" validate:"required"`
	Breakaway   string       `json:"breakaway" validate:"required"`
}

type FetchUpdatesWhereInputDto struct {
	EmployeeID *string    `json:"employee_id" `
	SprintId   *string    `json:"sprint_id" validate:"required"`
	WeekNumber *int       `json:"week_number"`
	DayNumber  *int       `json:"day_number"`
	// DateBefore *dto.JsonDate `json:"date_before"`
	// DateAfter  *dto.JsonDate `json:"date_after"`
}

type UpdateDto struct {
	ID           string       `json:"id"`
	EmployeeID   string       `json:"employee_id"`
	// EmployeeName string       `json:"employee_name"`
	SprintID     string       `json:"sprint_id"`
	CreatedAt    time.Time    `json:"created_at"`
	Yesterday    string       `json:"yesterday"`
	Today        string       `json:"today"`
	BlockedBy    []string     `json:"blocked_by"`
	Breakaway    string       `json:"breakaway"`
	CheckInTime  dto.Json24HrTime    `json:"check_in_time"`
	Status       UpdateStatus `json:"status"`
}

type UpdateStatus string

const (
	StatusBefore UpdateStatus = "before standup";
	StatusAfter UpdateStatus = "after standup";
	StatusWithin UpdateStatus = "within standup"
)
