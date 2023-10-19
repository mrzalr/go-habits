package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/internal/formatter"
)

func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return c.Next()
	}

	traceID := uuid.New().String()
	c.Set("X-Trace-ID", traceID)

	return formatter.SendErrorResponse(c, err, traceID)
}
