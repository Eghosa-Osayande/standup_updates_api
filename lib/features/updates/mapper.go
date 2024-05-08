package updates

import (
	"standup-api/lib/common/database"
)

func UpdateModelToDto(m *database.Update) *UpdateDto {
	var id, _ = m.ID.Value()

	employeeId, _ := m.EmployeeID.Value()
	sprintId, _ := m.SprintID.Value()

	return &UpdateDto{
		ID:         id.(string),
		EmployeeID: employeeId.(string),
		SprintID:   sprintId.(string),
		CreatedAt:  m.CreatedAt.Time,
		Yesterday:  m.Yesterday,
		Today:      m.Today,
		BlockedBy:  m.BlockedBy,
		Breakaway:  m.Breakaway,
		Status:     UpdateStatus(m.Status),
	}
}

func FetchAllUpdatesRowToDto(m *database.FetchAllUpdatesRow) *UpdateDto {
	id, _ := m.ID.Value()
	employeeId, _ := m.EmployeeID.Value()
	sprintId, _ := m.SprintID.Value()

	return &UpdateDto{
		ID:           id.(string),
		EmployeeID:   employeeId.(string),
		EmployeeName: m.EmployeeName,
		SprintID:     sprintId.(string),
		CreatedAt:    m.CreatedAt.Time,
		Yesterday:    m.Yesterday,
		Today:        m.Today,
		BlockedBy:    m.BlockedBy,
		Breakaway:    m.Breakaway,
		Status:       UpdateStatus(m.Status),
	}
}
