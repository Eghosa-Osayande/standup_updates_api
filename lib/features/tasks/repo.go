package sprint


type TasksRepository interface {
	CreateTask(*CreateTaskInputDto) (*TasksDto, error)
}