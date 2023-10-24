package mysql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/internal/habit/repository/mysql/query"
	"github.com/mrzalr/go-habits/pkg/date"
)

func (r *repository) GetHabits(weekRange date.WeekRange) ([]model.HabitResponse, error) {
	nstmt, err := r.db.PrepareNamed(query.Habit.GetAllHabits())
	if err != nil {
		return nil, err
	}
	defer nstmt.Close()

	rows, err := nstmt.Query(weekRange)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDataNotFound
		}
		return nil, err
	}
	defer rows.Close()

	habits := []model.HabitResponse{}
	for rows.Next() {
		habit := model.HabitResponse{}
		err := rows.Scan(
			&habit.ID, &habit.Category, &habit.Activity, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}

	return habits, nil
}

func (r *repository) GetHabitByID(id uuid.UUID) (model.HabitResponse, error) {
	nstmt, err := r.db.PrepareNamed(query.Habit.GetHabitByID())
	if err != nil {
		return model.HabitResponse{}, err
	}
	defer nstmt.Close()

	params := queryParams{"id": id}

	habit := model.HabitResponse{}
	err = nstmt.QueryRow(params).Scan(
		&habit.ID, &habit.Category, &habit.Activity, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.HabitResponse{}, model.ErrDataNotFound
		}
		return model.HabitResponse{}, err
	}

	return habit, nil
}

func (r *repository) CreateHabit(habit model.Habit) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(query.Habit.CreateHabit())
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(habit)
	if err != nil {
		return uuid.UUID{}, err
	}

	return habit.ID, nil
}

func (r *repository) UpdateHabit(habit model.Habit) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(query.Habit.UpdateHabit())
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(habit)
	if err != nil {
		return uuid.UUID{}, err
	}

	return habit.ID, nil
}
