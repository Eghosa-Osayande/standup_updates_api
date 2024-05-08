package updates

import (
	"standup-api/lib/common/database"
	"time"
)


func UpdateModelToDto(m *database.Update) *UpdateDto {
	var id, _ = m.ID.Value()
	createdAt, _ := m.CreatedAt.Value()
	employeeId,_:=m.EmployeeID.Value()
	sprintId,_:=m.SprintID.Value()
	checkInTime,_:=m.CheckInTime.Value()

	return &UpdateDto{
		ID:           id.(string),
		EmployeeID:   employeeId.(string),
		SprintID:     sprintId.(string),
		CreatedAt:    createdAt.(time.Time),
		Yesterday:    m.Yesterday,
		Today:        m.Today,
		BlockedBy:    m.BlockedBy,
		Breakaway:    m.Breakaway,
		CheckInTime:  checkInTime.(time.Time),
		Status:       UpdateStatus(m.Status),
	}
}