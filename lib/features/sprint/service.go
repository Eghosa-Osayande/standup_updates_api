package sprint

import (
	"errors"
	"standup-api/lib/utils/http_response"
)


type SprintService interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	UpdateSprint(*UpdateSprintInputDto) (*SprintDto, error)
	StartSprint() (*SprintDto, *http_response.HttpError)
}

func NewSprintService(repo SprintsRepository) SprintService {
	return &sprintService{sprintRepo: repo}
}

type sprintService struct {
	sprintRepo SprintsRepository
}

func (s *sprintService) CreateSprint(input *CreateSprintInputDto) (*SprintDto, error) {
	displayErr := errors.New("error creating sprint")

	sprint, err := s.sprintRepo.CreateSprint(input)
	if err != nil {
		return nil, displayErr
	}

	return sprint, nil
}

func (s *sprintService) UpdateSprint(input *UpdateSprintInputDto) (*SprintDto, error) {
	displayErr := errors.New("error updating sprint")

	sprint, err := s.sprintRepo.UpdateSprint(input)
	if err != nil {
		return nil, displayErr
	}

	return sprint, nil
}

func (s *sprintService) StartSprint() (*SprintDto, *http_response.HttpError) {
	// displayErr := errors.New("error starting sprint")

	// sprint, err := s.sprintRepo.StartSprint()
	// if err != nil {
	// 	return nil, displayErr
	// }

	return nil, http_response.NewHttpError(409,"sprint already started")
}
