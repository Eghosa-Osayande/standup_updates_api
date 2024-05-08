package sprint

import (
	"context"
	"standup-api/lib/common/database"
	"standup-api/lib/utils/http_response"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SprintsRepository interface {
	CreateSprint(*CreateSprintInputDto) (*SprintDto, error)
	FetchAllSprints(input *FetchSprintsInputDto) (*http_response.CursorPage[SprintDto], error)
}

func NewSprintRepo(db *database.Database) SprintsRepository {
	return &sprintRepo{
		db: db,
	}
}

type sprintRepo struct {
	db *database.Database
}

func (r *sprintRepo) CreateSprint(input *CreateSprintInputDto) (*SprintDto, error) {
	sprint,err:=r.db.CreateSprint(context.Background(), database.CreateSprintParams{
		Name:             "",
		StandupStartTime: pgtype.Timestamptz{
			Time:input.StandupStartTime.ToTime(),
			Valid: true,
		},
		StandupEndTime:   pgtype.Timestamptz{
			Time:input.StandupEndTime.ToTime(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return SprintModelToDto(&sprint), nil
}


func (r *sprintRepo) FetchAllSprints(input *FetchSprintsInputDto) (*http_response.CursorPage[SprintDto], error){

	sprints, err := r.db.FetchAllSprint(context.Background(), database.FetchAllSprintParams{
		Offset: pgtype.Int4{Int32: int32((input.Page - 1) * input.PerPage), Valid: true},
		Limit:  pgtype.Int4{Int32: int32(input.PerPage), Valid: true},
	})
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	var sprintSlice []SprintDto = make([]SprintDto, 0)

	for _, sprintModel := range sprints {
		sprintSlice = append(sprintSlice, *SprintModelToDto(&sprintModel))

	}

	return &http_response.CursorPage[SprintDto]{Data: sprintSlice, Cursor: &http_response.Cursor{NextPage: input.Page + 1, PerPage: input.PerPage}}, nil
}