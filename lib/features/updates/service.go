package updates

import "standup-api/lib/utils/pagination"

type UpdateService interface {
	CreateUpdate(*CreateUpdateInputDto) (*UpdateDto, error)
	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*pagination.CursorPage[UpdateDto], error)
}
