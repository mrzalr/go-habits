package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/formatter"
	"github.com/mrzalr/go-habits/internal/habit/model"
	_validator "github.com/mrzalr/go-habits/pkg/validator"
)

func (h *handler) GetHabits(c *fiber.Ctx) error {
	result, err := h.usecase.GetHabits()
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}

func (h *handler) CreateHabit(c *fiber.Ctx) error {
	payload := model.NewHabit()
	if err := c.BodyParser(&payload); err != nil {
		return model.ErrBadRequest
	}

	err := _validator.ValidateStruct(payload)
	if err != nil {
		return err
	}

	result, err := h.usecase.CreateHabit(payload)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusCreated, result)
}

func (h *handler) UpdateHabit(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return model.ErrInvalidID
	}

	payload := model.Habit{}
	if err := c.BodyParser(&payload); err != nil {
		return model.ErrBadRequest
	}

	result, err := h.usecase.UpdateHabit(id, payload)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}
