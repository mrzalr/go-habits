package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/formatter"
	"github.com/mrzalr/go-habits/internal/habit"
	"github.com/mrzalr/go-habits/internal/habit/model"
	_validator "github.com/mrzalr/go-habits/pkg/validator"
)

type handler struct {
	usecase habit.Usecase
}

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

func (h *handler) StartActivity(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))

	result, err := h.usecase.StartActivity(id)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}

func (h *handler) StopActivity(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))

	result, err := h.usecase.StopActivity(id)
	if err != nil {
		return err
	}

	return formatter.SendSuccessResponse(c, common.StatusOk, result)
}

func New(app *fiber.App, usecase habit.Usecase) {
	h := handler{
		usecase: usecase,
	}

	v1 := app.Group("/v1")
	v1.Get("/habits", h.GetHabits)
	v1.Post("/habits", h.CreateHabit)
	v1.Patch("/habits/:id/start", h.StartActivity)
	v1.Patch("/habits/:id/done", h.StopActivity)
}
