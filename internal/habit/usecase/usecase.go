package usecase

import (
	"time"

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
	return u.repository.CreateHabit(habit)
}

func New(repository habit.Repository) habit.Usecase {
	return &usecase{
		repository: repository,
	}
}
