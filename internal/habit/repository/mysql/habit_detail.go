package mysql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

func (r *repository) GetHabitDetailByID(id uuid.UUID) (model.HabitDetail, error) {
	nstmt, err := r.db.PrepareNamed(GetHabitDetailByIDQuery)
	if err != nil {
		return model.HabitDetail{}, err
	}
	defer nstmt.Close()

	params := queryParams{"id": id}

	habitDetail := model.HabitDetail{}
	err = nstmt.QueryRow(params).Scan(
		&habitDetail.ID,
		&habitDetail.HabitID,
		&habitDetail.StartTime,
		&habitDetail.EndTime,
		&habitDetail.Remark,
		&habitDetail.Valid,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.HabitDetail{}, model.ErrDataNotFound
		}
		return model.HabitDetail{}, err
	}

	return habitDetail, nil
}

func (r *repository) CreateHabitDetail(habitDetail model.HabitDetail) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(CreateHabitDetailQuery)
	if err != nil {
		return uuid.UUID{}, nil
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(habitDetail)
	if err != nil {
		return uuid.UUID{}, nil
	}

	return habitDetail.ID, nil
}

func (r *repository) GetLastHabitDetailStarted(habitID uuid.UUID) (model.HabitDetail, error) {
	nstmt, err := r.db.PrepareNamed(GetLastHabitDetailStartedQuery)
	if err != nil {
		return model.HabitDetail{}, err
	}
	defer nstmt.Close()

	params := queryParams{"habit_id": habitID, "end_time": time.Time{}}

	habitDetail := model.HabitDetail{}
	err = nstmt.QueryRow(params).Scan(
		&habitDetail.ID,
		&habitDetail.HabitID,
		&habitDetail.StartTime,
		&habitDetail.EndTime,
		&habitDetail.Remark,
		&habitDetail.Valid,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.HabitDetail{}, model.ErrDataNotFound
		}
		return model.HabitDetail{}, err
	}

	return habitDetail, nil
}

func (r *repository) UpdateHabitDetail(habitDetail model.HabitDetail) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(UpdateHabitDetailQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(habitDetail)
	if err != nil {
		return uuid.UUID{}, err
	}

	return habitDetail.ID, nil
}
