package sprint

type SprintsRepository interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	UpdateSprint(*UpdateSprintInputDto) (*SprintDto, error)
}
