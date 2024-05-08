package sprint

import (
	"standup-api/lib/common/database"
)

func SprintModelToDto(m *database.Sprint) *SprintDto {
	var id, _ = m.ID.Value()

	startTimeFormatted:= Json24HrTime(m.StandupStartTime.Time.UTC())
	endTimeFormatted:= Json24HrTime(m.StandupEndTime.Time.UTC())
	
	return &SprintDto{
		ID:               id.(string),
		Name:             m.Name,
		CreatedAt:        m.CreatedAt.Time,
		StandupStartTime: &startTimeFormatted,
		StandupEndTime:   &endTimeFormatted,
	}
}
