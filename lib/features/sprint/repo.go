package sprint

import "standup-api/lib/utils/http_response"

type SprintsRepository interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	UpdateSprint(*UpdateSprintInputDto) (*SprintDto, error)
	FetchAllSprints(input *FetchSprintsInputDto) (*http_response.CursorPage[SprintDto], error)
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

func (r *sprintRepo) FetchAllSprints(input *FetchSprintsInputDto) (*http_response.CursorPage[SprintDto], error){

	return &http_response.CursorPage[SprintDto]{},nil
}