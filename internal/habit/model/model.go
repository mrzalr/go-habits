package model

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Activity    string    `json:"activity" db:"activity" validate:"required,min=8"`
	Description string    `json:"description" db:"description"`
	StartTime   time.Time `json:"start_time" db:"startTime"`
	EndTime     time.Time `json:"end_time" db:"endTime"`
	CreatedAt   time.Time `json:"created_at" db:"createdAt"`
}

func (h *Habit) Start() error {
	if h.StartTime != (time.Time{}) {
		return ErrAlreadyStarted
	}

	h.StartTime = time.Now()
	return nil
}

func (h *Habit) Stop() {
	h.EndTime = time.Now()
}

func NewHabit() Habit {
	return Habit{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
