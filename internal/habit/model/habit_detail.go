package model

import (
	"time"

	"github.com/google/uuid"
)

type HabitDetail struct {
	ID        uuid.UUID `json:"id" db:"id"`
	HabitID   uuid.UUID `json:"habit_id" db:"habitID"`
	StartTime time.Time `json:"start_time" db:"startTime"`
	EndTime   time.Time `json:"end_time" db:"endTime"`
	Remark    string    `json:"remark" db:"remark"`
	Valid     bool      `json:"valid" db:"valid"`
}

func NewHabitDetail(habitID uuid.UUID) HabitDetail {
	return HabitDetail{
		ID:        uuid.New(),
		HabitID:   habitID,
		StartTime: time.Now(),
		Valid:     false,
	}
}

func (hd *HabitDetail) Stop(remark string) {
	hd.EndTime = time.Now()
	hd.Remark = remark
	hd.Valid = true
}

type StopHabitRequest struct {
	Remark string `json:"remark" validate:"required"`
}
