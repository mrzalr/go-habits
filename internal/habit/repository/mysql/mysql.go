package mysql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mrzalr/go-habits/internal/habit"
	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

type queryParams map[string]any

type repository struct {
	db *sqlx.DB
}

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
		m_habit := model.Habit{}
		err := rows.Scan(
			&m_habit.ID, &m_habit.Activity, &m_habit.Description, &m_habit.StartTime, &m_habit.EndTime, &m_habit.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		habits = append(habits, m_habit)
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

	m_habit := model.Habit{}
	err = nstmt.QueryRow(params).Scan(
		&m_habit.ID, &m_habit.Activity, &m_habit.Description, &m_habit.StartTime, &m_habit.EndTime, &m_habit.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Habit{}, model.ErrDataNotFound
		}
		return model.Habit{}, err
	}

	return m_habit, nil
}

func (r *repository) CreateHabit(m_habit model.Habit) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(CreateHabitQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(m_habit)
	if err != nil {
		return uuid.UUID{}, err
	}

	return m_habit.ID, nil
}

func (r *repository) UpdateHabit(id uuid.UUID, m_habit model.Habit) (uuid.UUID, error) {
	nstmt, err := r.db.PrepareNamed(UpdateHabitQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(m_habit)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func New(db *sqlx.DB) habit.Repository {
	return &repository{
		db: db,
	}
}
