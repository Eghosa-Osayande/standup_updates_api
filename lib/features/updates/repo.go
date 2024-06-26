package updates

import (
	"context"
	"log"
	"standup-api/lib/common/database"
	"standup-api/lib/utils/http_response"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdatesRepository interface {
	CreateUpdate(input *CreateUpdateInputDto) (*UpdateDto, error)

	FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response.CursorPage[UpdateDto], error)
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

	employeeId := pgtype.UUID{}
	employeeId.Scan(input.EmployeeID)
	sprintId := pgtype.UUID{}
	sprintId.Scan(input.SprintID)
	sprint, err := r.db.FetchSprintById(context.Background(), sprintId)

	if err != nil {
		return nil, err
	}

	timeToDuration := func(t time.Time) time.Duration {
		return time.Duration(t.Hour()*int(time.Hour) + t.Minute()*int(time.Minute))
	}
	checkedInAt := time.Now().UTC()
	log.Println(checkedInAt)
	checkInDuration := timeToDuration(checkedInAt)
	start := timeToDuration(sprint.StandupStartTime.Time.UTC())
	end := timeToDuration(sprint.StandupEndTime.Time.UTC())

	var status UpdateStatus

	if checkInDuration >= start && checkInDuration <= end {
		status = StatusWithin
	} else if checkInDuration < start {
		status = StatusBefore
	} else {
		status = StatusAfter
	}

	update, err := r.db.CreateUpdate(context.Background(), database.CreateUpdateParams{
		EmployeeID: employeeId,
		SprintID:   sprintId,
		Yesterday:  input.Yesterday,
		Today:      input.Today,
		BlockedBy:  input.BlockedBy,
		Breakaway:  input.Breakaway,
		CheckInTime: pgtype.Timestamptz{
			Time:  checkedInAt,
			Valid: true,
		},
		Status: string(status),
		Tasks:  input.Tasks,
	})
	if err != nil {
		return nil, err
	}
	return UpdateModelToDto(&update), nil
}

func (r *updatesRepo) FindUpdatesWhere(input *FetchUpdatesWhereInputDto) (*http_response.CursorPage[UpdateDto], error) {
	start := pgtype.Timestamptz{}
	end := pgtype.Timestamptz{}

	if input.StartDate != nil {
		t := input.EndDate.ToTime()
		year, month, day := t.Date()
		newTime := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		start.Scan(newTime)
	}
	
	if input.EndDate != nil {
		t := input.EndDate.ToTime()
		year, month, day := t.Date()
		newTime := time.Date(year, month, day, 23, 59, 59, 0, t.Location())
		end.Scan(newTime)
	}
	updates, err := r.db.FetchAllUpdates(context.Background(), database.FetchAllUpdatesParams{
		SprintID:   input.SprintId,
		EmployeeID: input.EmployeeID,
		StartDate:  start,
		EndDate:    end,
		Offset:     pgtype.Int4{Int32: int32((input.Page - 1) * input.PerPage), Valid: true},
		Limit:      pgtype.Int4{Int32: int32(input.PerPage), Valid: true},
	})
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	var sprintSlice []UpdateDto = make([]UpdateDto, 0)

	for _, updateModel := range updates {
		sprintSlice = append(sprintSlice, *FetchAllUpdatesRowToDto(&updateModel))

	}

	return &http_response.CursorPage[UpdateDto]{Data: sprintSlice, Cursor: &http_response.Cursor{NextPage: input.Page + 1, PerPage: input.PerPage}}, nil

}
