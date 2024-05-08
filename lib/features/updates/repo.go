package updates

import "standup-api/lib/utils/pagination"



type UpdatesRepository interface {
	CreateUpdate(*CreateUpdateInputDto) (*UpdateDto, error)
	
	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*pagination .CursorPage[UpdateDto], error)
}