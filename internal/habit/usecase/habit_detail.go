package usecase

import (
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

func (u *usecase) StartActivity(habitID uuid.UUID) (model.HabitDetail, error) {
	lastHabitStarted, err := u.repository.GetLastHabitDetailStarted(habitID)
	if err != nil && err != model.ErrDataNotFound {
		return model.HabitDetail{}, err
	}

	if lastHabitStarted != (model.HabitDetail{}) {
		return model.HabitDetail{}, model.ErrAlreadyStarted
	}

	habitDetail := model.NewHabitDetail(habitID)

	insertedID, err := u.repository.CreateHabitDetail(habitDetail)
	if err != nil {
		return model.HabitDetail{}, err
	}

	return u.repository.GetHabitDetailByID(insertedID)
}

func (u *usecase) StopActivity(habitID uuid.UUID, stopHabitRequest model.StopHabitRequest) (model.HabitDetail, error) {
	lastHabitStarted, err := u.repository.GetLastHabitDetailStarted(habitID)
	if err != nil {
		return model.HabitDetail{}, err
	}

	lastHabitStarted.Stop(stopHabitRequest.Remark)

	updatedID, err := u.repository.UpdateHabitDetail(lastHabitStarted)
	if err != nil {
		return model.HabitDetail{}, err
	}

	return u.repository.GetHabitDetailByID(updatedID)
}
