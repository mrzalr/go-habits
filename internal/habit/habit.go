package habit

import (
	"github.com/google/uuid"

	"github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/date"
)

type Repository interface {
	GetHabitCategories() ([]model.Category, error)
	GetHabitCategoryByID(id uuid.UUID) (model.Category, error)
	CreateHabitCategory(category model.Category) (uuid.UUID, error)
	UpdateHabitCategory(category model.Category) (uuid.UUID, error)

	GetHabits(weekRange date.WeekRange) ([]model.Habit, error)
	CreateHabit(habit model.Habit) (uuid.UUID, error)
	GetHabitByID(id uuid.UUID) (model.Habit, error)

	GetHabitDetailByID(id uuid.UUID) (model.HabitDetail, error)
	GetLastHabitDetailStarted(habitID uuid.UUID) (model.HabitDetail, error)
	CreateHabitDetail(habitDetail model.HabitDetail) (uuid.UUID, error)
	UpdateHabitDetail(habitDetail model.HabitDetail) (uuid.UUID, error)
}

type Usecase interface {
	GetHabitCategories() ([]model.Category, error)
	CreateHabitCategory(category model.Category) (model.Category, error)
	UpdateHabitCategory(id uuid.UUID, category model.Category) (model.Category, error)

	GetHabits() ([]model.Habit, error)
	CreateHabit(habit model.Habit) (model.Habit, error)

	StartActivity(id uuid.UUID) (model.HabitDetail, error)
	StopActivity(id uuid.UUID, remark model.StopHabitRequest) (model.HabitDetail, error)
}
