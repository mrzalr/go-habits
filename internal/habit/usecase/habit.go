package usecase

import (
	"time"

	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

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
