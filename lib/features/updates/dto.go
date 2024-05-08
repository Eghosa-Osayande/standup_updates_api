package updates

import "time"

type CreateUpdateInputDto struct {
	EmployeeID  string       `json:"employee_id"`
	SprintID    string       `json:"sprint_id"`
	Yesterday   string       `json:"yesterday"`
	Today       string       `json:"today"`
	BlockedBy   []string     `json:"blocked_by"`
	Breakaway   string       `json:"breakaway"`
	CheckInTime time.Time    `json:"check_in_time"`
	Status      UpdateStatus `json:"status"`
}

type FetchUpdatesWhereInputDto struct {
	EmployeeID *string    `json:"employee_id"`
	SprintId   *string    `json:"sprint_id"`
	WeekNumber *int       `json:"week_number"`
	DayNumber  *int       `json:"day_number"`
	DateBefore *time.Time `json:"date_before"`
	DateAfter  *time.Time `json:"date_after"`
}

type UpdateDto struct {
	ID           string       `json:"id"`
	EmployeeID   string       `json:"employee_id"`
	EmployeeName string       `json:"employee_name"`
	SprintID     string       `json:"sprint_id"`
	CreatedAt    time.Time    `json:"created_at"`
	Yesterday    string       `json:"yesterday"`
	Today        string       `json:"today"`
	BlockedBy    []string     `json:"blocked_by"`
	Breakaway    string       `json:"breakaway"`
	CheckInTime  time.Time    `json:"check_in_time"`
	Status       UpdateStatus `json:"status"`
}

type UpdateStatus int

const (
	StatusBefore UpdateStatus = iota
	StatusAfter
	StatusWithin
)
