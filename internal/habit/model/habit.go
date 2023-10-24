package model

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CategoryID  uuid.UUID `json:"category_id" db:"categoryID" validate:"required"`
	Activity    string    `json:"activity" db:"activity" validate:"required,min=8"`
	Description string    `json:"description" db:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at" db:"updatedAt"`
}

func NewHabit() Habit {
	return Habit{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
