package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/formatter"
	"github.com/mrzalr/go-habits/internal/habit/model"
	_validator "github.com/mrzalr/go-habits/pkg/validator"
)

func (h *handler) GetHabitCategories(c *fiber.Ctx) error {
	result, err := h.usecase.GetHabitCategories()
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}

func (h *handler) CreateHabitCategory(c *fiber.Ctx) error {
	payload := model.NewCategory()
	if err := c.BodyParser(&payload); err != nil {
		return model.ErrBadRequest
	}

	err := _validator.ValidateStruct(payload)
	if err != nil {
		return err
	}

	result, err := h.usecase.CreateHabitCategory(payload)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusCreated, result)
}

func (h *handler) UpdateHabitCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return model.ErrInvalidID
	}

	payload := model.Category{UpdatedAt: time.Now()}
	if err := c.BodyParser(&payload); err != nil {
		return model.ErrBadRequest
	}

	err = _validator.ValidateStruct(payload)
	if err != nil {
		return err
	}

	result, err := h.usecase.UpdateHabitCategory(id, payload)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}
