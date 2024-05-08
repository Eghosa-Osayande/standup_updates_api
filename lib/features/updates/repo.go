package updates

import (
	"standup-api/lib/common/database"
	"standup-api/lib/utils/http_response"
)



type UpdatesRepository interface {
	CreateUpdate(*CreateUpdateInputDto) (*UpdateDto, error)
	
	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response .CursorPage[UpdateDto], error)
}

func NewUpdatesRepo(db *database.Database) UpdatesRepository {
	return &updatesRepo{
		db: db,
	}
}

type updatesRepo struct {
	db *database.Database
}

func (r *updatesRepo) CreateUpdate(input *CreateUpdateInputDto) (*UpdateDto, error) {
	return &UpdateDto{}, nil
}

func (r *updatesRepo) FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response.CursorPage[UpdateDto], error) {
	return &http_response.CursorPage[UpdateDto]{}, nil
}