package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal"
	"github.com/mrzalr/go-habits/internal/welcome"
)

type handler struct {
	usecase welcome.Usecase
}

func New(app *fiber.App, usecase welcome.Usecase) {
	h := handler{
		usecase: usecase,
	}

	v1 := app.Group("v1")
	v1.Get("/welcome", h.GetMessage)
}

func (h *handler) GetMessage(c *fiber.Ctx) error {
	name := c.Query("name")

	welcome, err := h.usecase.GetMessage(name)
	if err != nil {
		//handle error
	}

	return c.JSON(internal.NewResponseOK(welcome))
}
