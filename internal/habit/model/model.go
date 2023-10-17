package model

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Activity    string    `json:"activity" db:"activity"`
	Description string    `json:"description" db:"description"`
	StartTime   time.Time `json:"start_time" db:"startTime"`
	EndTime     time.Time `json:"end_time" db:"endTime"`
	CreatedAt   time.Time `json:"created_at" db:"createdAt"`
}

func NewHabit() Habit {
	return Habit{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
