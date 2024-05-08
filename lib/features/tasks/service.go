package sprint

type TaskService interface {
	CreateTask(*CreateTaskInputDto) (*TasksDto, error)
}
