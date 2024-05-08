package sprint


type SprintService interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	UpdateSprint(*UpdateSprintInputDto) (*SprintDto, error)
}
