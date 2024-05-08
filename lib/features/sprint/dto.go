package sprint

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type CreateSprintInputDto struct {
	Name             string    `json:"name" validate:"required"`
	StandupStartTime *Json24HrTime `json:"standup_start_time" validate:"required"`
	StandupEndTime   *Json24HrTime `json:"standup_end_time" validate:"required"`
}



type SprintDto struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"created_at"`
	StandupStartTime *Json24HrTime `json:"standup_start_time"`
	StandupEndTime   *Json24HrTime `json:"standup_end_time"`
}

type FetchSprintsInputDto struct {
	Page int `json:"page" validate:"required"`
	PerPage int `json:"per_page" validate:"required"`
}


type Json24HrTime time.Time

var hrs24Format = "15:04"

func (j *Json24HrTime) UnmarshalJSON(b []byte) error {
	
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(hrs24Format, s,time.UTC)
	if err != nil {
		return fmt.Errorf("error parsing date %s, required format %s", s, hrs24Format)
	}

	*j = Json24HrTime(t)
	return nil
}

func (j Json24HrTime) MarshalJSON() ([]byte, error) {
	timeString := time.Time(j).Format(hrs24Format)
	return json.Marshal(timeString + " UTC")
}

func (j *Json24HrTime) ToTime() time.Time {
	return time.Time(*j)
}
