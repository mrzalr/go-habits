package mysql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

func (r *repository) GetHabits(weekRange date.WeekRange) ([]model.Habit, error) {
	nstmt, err := r.db.PrepareNamed(GetAllHabitsQuery)
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

	habits := []model.Habit{}
	for rows.Next() {
		habit := model.Habit{}
		err := rows.Scan(
			&habit.ID, &habit.CategoryID, &habit.Activity, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}

	return habits, nil
}

func (r *repository) GetHabitByID(id uuid.UUID) (model.Habit, error) {
	nstmt, err := r.db.PrepareNamed(GetHabitByIDQuery)
	if err != nil {
		return model.Habit{}, err
	}
	defer nstmt.Close()

	params := queryParams{"id": id}

	habit := model.Habit{}
	err = nstmt.QueryRow(params).Scan(
		&habit.ID, &habit.CategoryID, &habit.Activity, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Habit{}, model.ErrDataNotFound
		}
		return model.Habit{}, err
	}

	return habit, nil
}

func (r *repository) CreateHabit(habit model.Habit) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(CreateHabitQuery)
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
	nstmt, err := r.db.PrepareNamed(UpdateHabitQuery)
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
