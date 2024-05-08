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
