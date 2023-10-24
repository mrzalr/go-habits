package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

func (u *usecase) GetHabits() ([]model.HabitResponse, error) {
	today := time.Now()
	wd := int(today.Weekday())
	weekRange := date.GetWeekRange(wd)

	habits, err := u.repository.GetHabits(weekRange)
	if err != nil {
		return nil, err
	}

	return habits, nil
}

func (u *usecase) CreateHabit(habit model.Habit) (model.HabitResponse, error) {
	insertedID, err := u.repository.CreateHabit(habit)
	if err != nil {
		return model.HabitResponse{}, err
	}

	return u.repository.GetHabitByID(insertedID)
}

func (u *usecase) UpdateHabit(id uuid.UUID, habit model.Habit) (model.HabitResponse, error) {
	foundHabit, err := u.repository.GetHabitByID(id)
	if err != nil {
		return model.HabitResponse{}, err
	}

	habit.ID = foundHabit.ID
	habit.UpdatedAt = time.Now()

	updatedID, err := u.repository.UpdateHabit(habit)
	if err != nil {
		return model.HabitResponse{}, err
	}

	return u.repository.GetHabitByID(updatedID)
}
