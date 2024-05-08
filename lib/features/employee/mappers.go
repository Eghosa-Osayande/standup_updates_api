package employee

import "standup-api/lib/common/database"

func EmployeeModelToDto(m *database.Employee) *EmployeeDto {
	var id, _ = m.ID.Value()
	return &EmployeeDto{
		ID:        id.(string),
		Name:      m.Name,
		Password:  m.Password,
		CreatedAt: m.CreatedAt.Time,
		Role:      m.Role,
	}
}
