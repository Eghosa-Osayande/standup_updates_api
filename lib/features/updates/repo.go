package updates

import "standup-api/lib/utils/pagination"



type UpdatesRepository interface {
	CreateUpdate(*CreateUpdateInputDto) (*UpdateDto, error)
	
	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*pagination .CursorPage[UpdateDto], error)
}

func NewUpdatesRepo() UpdatesRepository {
	return &updatesRepo{}
}

type updatesRepo struct {
}

func (r *updatesRepo) CreateUpdate(input *CreateUpdateInputDto) (*UpdateDto, error) {
	return &UpdateDto{}, nil
}

func (r *updatesRepo) FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*pagination.CursorPage[UpdateDto], error) {
	return &pagination.CursorPage[UpdateDto]{}, nil
}