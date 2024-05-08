package dto

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)


type JsonDate time.Time

var dateFormat = "2006-01-02"

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(dateFormat, s, time.UTC)
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

type Json24HrTime time.Time

var hrs24Format = "15:04"

func (j *Json24HrTime) UnmarshalJSON(b []byte) error {
	
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(hrs24Format, s, time.UTC)
	if err != nil {
		return fmt.Errorf("error parsing date %s, required format %s", s, hrs24Format)
	}

	*j = Json24HrTime(t)
	return nil
}

func (j Json24HrTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j).Format(hrs24Format))
}

func (j *Json24HrTime) ToTime() time.Time {
	return time.Time(*j)
}
