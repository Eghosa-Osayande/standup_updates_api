package sprint

import (
	"standup-api/lib/common/database"
	"standup-api/lib/common/dto"
)

func SprintModelToDto(m *database.Sprint) *SprintDto {
	var id, _ = m.ID.Value()

	startTimeFormatted:= dto.Json24HrTime(m.StandupStartTime.Time)
	endTimeFormatted:= dto.Json24HrTime(m.StandupEndTime.Time)
	
	return &SprintDto{
		ID:               id.(string),
		Name:             m.Name,
		CreatedAt:        m.CreatedAt.Time,
		StandupStartTime: &startTimeFormatted,
		StandupEndTime:   &endTimeFormatted,
	}
}
