package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/formatter"
	"github.com/mrzalr/go-habits/internal/habit"
	"github.com/mrzalr/go-habits/internal/habit/model"
)

type handler struct {
	usecase habit.Usecase
}

func (h *handler) GetHabits(c *fiber.Ctx) error {
	result, err := h.usecase.GetHabits()
	if err != nil {
		return err
	}

	return formatter.NewResponseOk(c, result)
}

func (h *handler) CreateHabit(c *fiber.Ctx) error {
	payload := model.NewHabit()
	if err := c.BodyParser(&payload); err != nil {
		// TODO should return validate error
		return err
	}

	result, err := h.usecase.CreateHabit(payload)
	if err != nil {
		return err
	}

	return formatter.NewResponseCreated(c, result)
}

func New(app *fiber.App, usecase habit.Usecase) {
	h := handler{
		usecase: usecase,
	}

	v1 := app.Group("/v1")
	v1.Get("/habits", h.GetHabits)
	v1.Post("/habits", h.CreateHabit)
}
