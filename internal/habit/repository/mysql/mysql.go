package mysql

import (
	"database/sql"
	"log"

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
	query := `
	SELECT 
		id, activity, description, start_time, end_time, created_at 
	FROM habit
	WHERE created_at BETWEEN :startDate AND :endDate`

	nstmt, err := r.db.PrepareNamed(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer nstmt.Close()

	rows, err := nstmt.Query(weekRange)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, habit.ErrDataNotFound
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
			log.Println(err)
			return nil, err
		}
		habits = append(habits, m_habit)
	}

	return habits, err
}

func (r *repository) GetHabitByID(id uuid.UUID) (model.Habit, error) {
	query := `
	SELECT 
		id, activity, description, start_time, end_time, created_at 
	FROM habit
	WHERE id = :id`

	params := map[string]interface{}{"id": id}

	nstmt, err := r.db.PrepareNamed(query)
	if err != nil {
		log.Println(err)
		return model.Habit{}, err
	}
	defer nstmt.Close()

	m_habit := model.Habit{}
	err = nstmt.QueryRow(params).Scan(
		&m_habit.ID, &m_habit.Activity, &m_habit.Description, &m_habit.StartTime, &m_habit.EndTime, &m_habit.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return model.Habit{}, habit.ErrDataNotFound
		}
		return model.Habit{}, err
	}

	return m_habit, nil
}

func (r *repository) CreateHabit(m_habit model.Habit) (uuid.UUID, error) {
	query := `
	INSERT INTO 
		habit(id, activity, description, start_time, end_time, created_at)
	VALUES
		(:id, :activity, :description, :startTime, :endTime, :createdAt)`

	nstmt, err := r.db.PrepareNamed(query)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(m_habit)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, err
	}

	return m_habit.ID, nil
}

func (r *repository) UpdateHabit(id uuid.UUID, m_habit model.Habit) (uuid.UUID, error) {
	query := `
	UPDATE habit 
	SET 
		activity = :activity,
		description = :description,
		start_time = :startTime,
		end_time = :endTime
	WHERE id = :id
	`

	nstmt, err := r.db.PrepareNamed(query)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, err
	}
	defer nstmt.Close()

	_, err = nstmt.Exec(m_habit)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, err
	}

	return id, nil
}

func New(db *sqlx.DB) habit.Repository {
	return &repository{
		db: db,
	}
}
