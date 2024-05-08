package sprint

type SprintsRepository interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	UpdateSprint(*UpdateSprintInputDto) (*SprintDto, error)
}

func NewSprintRepo() SprintsRepository {
	return &sprintRepo{}
}

type sprintRepo struct {
}

func (r *sprintRepo) CreateSprint(input *CreateSprintInputDto) (*SprintDto, error) {
	return &SprintDto{}, nil
}

func (r *sprintRepo) UpdateSprint(input *UpdateSprintInputDto) (*SprintDto, error) {
	return &SprintDto{}, nil
}
