package sprint

import "time"



type CreateTaskInputDto struct {
	Name     string `json:"name"`
	SprintID string `json:"sprint_id"`
}


type TasksDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SprintId  string    `json:"sprint_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}