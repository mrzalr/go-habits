package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/formatter"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

func (h *handler) StartActivity(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return model.ErrInvalidID
	}

	result, err := h.usecase.StartActivity(id)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}

func (h *handler) StopActivity(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return model.ErrInvalidID
	}

	payload := model.StopHabitRequest{}
	if err := c.BodyParser(&payload); err != nil {
		return model.ErrBadRequest
	}

	result, err := h.usecase.StopActivity(id, payload)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}
