package sprint

import (
	"standup-api/lib/common/database"
	"time"
)

func SprintModelToDto(m *database.Sprint) *SprintDto {
	var id, _ = m.ID.Value()
	startTime, _ := m.StandupStartTime.Value()
	endTime, _ := m.StandupEndTime.Value()
	return &SprintDto{
		ID:               id.(string),
		Name:             m.Name,
		StandupStartTime: startTime.(time.Time),
		StandupEndTime:   endTime.(time.Time),
	}
}
