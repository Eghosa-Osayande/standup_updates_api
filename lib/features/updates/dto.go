package updates

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"strings"
	"time"
)

type CreateUpdateInputDto struct {
	EmployeeID string   `json:"employee_id" validate:"omitempty"`
	SprintID   string   `json:"sprint_id" validate:"required"`
	Yesterday  string   `json:"yesterday" validate:"required"`
	Today      string   `json:"today" validate:"required"`
	BlockedBy  []string `json:"blocked_by" validate:"required"`
	Tasks      []string `json:"tasks" validate:"required"`
	Breakaway  string   `json:"breakaway" validate:"required"`
}

type FetchUpdatesWhereInputDto struct {
	EmployeeID pgtype.UUID `json:"employee_id"`
	SprintId   pgtype.UUID `json:"sprint_id"`
	StartDate  *JsonDate   `json:"start_date"`
	EndDate    *JsonDate   `json:"end_date"`
	Page       int         `json:"page" validate:"required"`
	PerPage    int         `json:"per_page" validate:"required"`
}

type UpdateDto struct {
	ID         string `json:"id"`
	EmployeeID string `json:"employee_id"`
	EmployeeName string       `json:"employee_name,omitempty" `
	SprintID  string       `json:"sprint_id"`
	CreatedAt time.Time    `json:"created_at"`
	Yesterday string       `json:"yesterday"`
	Today     string       `json:"today"`
	BlockedBy []string     `json:"blocked_by"`
	Breakaway string       `json:"breakaway"`
	Status    UpdateStatus `json:"status"`
}

type UpdateStatus string

const (
	StatusBefore UpdateStatus = "before standup"
	StatusAfter  UpdateStatus = "after standup"
	StatusWithin UpdateStatus = "within standup"
)

// "2006-01-02T15:04:05"

type JsonDate time.Time

var dateFormat = "2006-01-02"

func (j *JsonDate) UnmarshalJSON(b []byte) error {

	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return fmt.Errorf("error parsing date %s, required format %s", s, dateFormat)
	}

	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j).Format(dateFormat))
}

func (j *JsonDate) ToTime() time.Time {
	return time.Time(*j)
}
