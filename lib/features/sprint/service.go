package sprint

import (
	"net/http"
	"standup-api/lib/utils/http_response"
)

type SprintService interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, *http_response.HttpError)

	FetchAllSprints(input *FetchSprintsInputDto) (*http_response.HttpPagedResponse[SprintDto], *http_response.HttpError)
}

func NewSprintService(repo SprintsRepository) SprintService {
	return &sprintService{sprintRepo: repo}
}

type sprintService struct {
	sprintRepo SprintsRepository
}

func (s *sprintService) CreateSprint(input *CreateSprintInputDto) (*SprintDto, *http_response.HttpError) {

	sprint, err := s.sprintRepo.CreateSprint(input)
	if err != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error creating sprint")
	}

	return sprint, nil
}



func (s *sprintService) FetchAllSprints(input *FetchSprintsInputDto) (*http_response.HttpPagedResponse[SprintDto], *http_response.HttpError) {
	sprints, err := s.sprintRepo.FetchAllSprints(input)
	if err != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error fetching sprints")
	}

	result := http_response.NewPagedResponse(*sprints)

	return &result, nil
}
