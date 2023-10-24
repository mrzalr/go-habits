package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/habit"
)

type handler struct {
	usecase habit.Usecase
}

func New(app *fiber.App, usecase habit.Usecase) {
	h := handler{
		usecase: usecase,
	}

	v1 := app.Group("/v1")

	v1.Get("/habits/categories", h.GetHabitCategories)
	v1.Post("/habits/categories", h.CreateHabitCategory)
	v1.Patch("/habits/categories/:id", h.UpdateHabitCategory)

	v1.Get("/habits", h.GetHabits)
	v1.Post("/habits", h.CreateHabit)

	v1.Patch("/habits/:id/start", h.StartActivity)
	v1.Patch("/habits/:id/end", h.StopActivity)
}
