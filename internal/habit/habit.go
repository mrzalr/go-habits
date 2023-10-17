package habit

import (
	"github.com/google/uuid"

	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

type Repository interface {
	GetHabits(weekRange date.WeekRange) ([]model.Habit, error)
	CreateHabit(habit model.Habit) (uuid.UUID, error)
	GetHabitByID(id uuid.UUID) (model.Habit, error)
	UpdateHabit(id uuid.UUID, habit model.Habit) (uuid.UUID, error)
}

type Usecase interface {
	GetHabits() ([]model.Habit, error)
	CreateHabit(habit model.Habit) (model.Habit, error)
	StartActivity(id uuid.UUID) (model.Habit, error)
	StopActivity(id uuid.UUID) (model.Habit, error)
}
