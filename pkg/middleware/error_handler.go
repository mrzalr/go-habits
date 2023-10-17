package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/habit"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := "internal server error"

	switch {
	case errors.Is(err, habit.ErrDataNotFound):
		status = fiber.StatusNotFound
		message = "not found"
	}

	c.Status(status)
	return c.JSON(common.NewErrorResponse(status, message, err.Error()))
}
