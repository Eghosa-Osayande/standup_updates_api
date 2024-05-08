package updates

import (
	"net/http"
	"standup-api/lib/utils/http_response"
)

type UpdateService interface {
	CreateUpdate(*CreateUpdateInputDto) (*http_response.HttpResponse[UpdateDto], *http_response.HttpError)

	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response.HttpResponse[http_response.CursorPage[UpdateDto]], error)
}

func NewUpdateService(updateRepo UpdatesRepository) UpdateService {
	return &updateService{updateRepo: updateRepo}
}

type updateService struct {
	updateRepo UpdatesRepository
}

func (s *updateService) CreateUpdate(input *CreateUpdateInputDto) (*http_response.HttpResponse[UpdateDto], *http_response.HttpError) {
	result, err := s.updateRepo.CreateUpdate(input)
	if err != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "failed to create update")
	}

	var response = http_response.NewSuccessResponse(*result)

	return &response, nil
}

func (s *updateService) FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response.HttpResponse[http_response.CursorPage[UpdateDto]], error) {
	result, err := s.updateRepo.FindUpdatesWhere(input)
	if err != nil {
		return nil, err
	}

	var response = http_response.NewSuccessResponse(*result)

	return &response, nil
}
