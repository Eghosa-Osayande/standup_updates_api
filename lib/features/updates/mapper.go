package updates

import (
	"standup-api/lib/common/database"
	"standup-api/lib/common/dto"
)

func UpdateModelToDto(m *database.Update) *UpdateDto {
	var id, _ = m.ID.Value()
	
	employeeId, _ := m.EmployeeID.Value()
	sprintId, _ := m.SprintID.Value()

	return &UpdateDto{
		ID:          id.(string),
		EmployeeID:  employeeId.(string),
		SprintID:    sprintId.(string),
		CreatedAt:   m.CreatedAt.Time,
		Yesterday:   m.Yesterday,
		Today:       m.Today,
		BlockedBy:   m.BlockedBy,
		Breakaway:   m.Breakaway,
		CheckInTime: dto.Json24HrTime(m.CheckInTime.Time.UTC()),
		Status:      UpdateStatus(m.Status),
	}
}
