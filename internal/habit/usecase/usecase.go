package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit"
	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

type usecase struct {
	repository habit.Repository
}

func (u *usecase) GetHabits() ([]model.Habit, error) {
	today := time.Now()
	wd := int(today.Weekday())
	weekRange := date.GetWeekRange(wd)

	habits, err := u.repository.GetHabits(weekRange)
	if err != nil {
		return nil, err
	}

	return habits, nil
}

func (u *usecase) CreateHabit(habit model.Habit) (model.Habit, error) {
	insertedID, err := u.repository.CreateHabit(habit)
	if err != nil {
		return model.Habit{}, err
	}

	return u.repository.GetHabitByID(insertedID)
}

func (u *usecase) StartActivity(id uuid.UUID) (model.Habit, error) {
	m_habit, err := u.repository.GetHabitByID(id)
	if err != nil {
		return model.Habit{}, err
	}

	if m_habit.StartTime != (time.Time{}) {
		return model.Habit{}, habit.ErrAlreadyStarted
	}
	m_habit.Start()

	updatedID, err := u.repository.UpdateHabit(id, m_habit)
	if err != nil {
		return model.Habit{}, err
	}

	return u.repository.GetHabitByID(updatedID)
}

func (u *usecase) StopActivity(id uuid.UUID) (model.Habit, error) {
	m_habit, err := u.repository.GetHabitByID(id)
	if err != nil {
		return model.Habit{}, err
	}

	m_habit.Stop()

	updatedID, err := u.repository.UpdateHabit(id, m_habit)
	if err != nil {
		return model.Habit{}, err
	}

	return u.repository.GetHabitByID(updatedID)
}

func New(repository habit.Repository) habit.Usecase {
	return &usecase{
		repository: repository,
	}
}
