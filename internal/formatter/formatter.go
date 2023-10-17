package formatter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/common"
)

func NewResponseOk(c *fiber.Ctx, data interface{}) error {
	c.Status(fiber.StatusOK)
	return c.JSON(common.NewSuccessResponse(fiber.StatusOK, "ok", data))
}

func NewResponseCreated(c *fiber.Ctx, data interface{}) error {
	c.Status(fiber.StatusCreated)
	return c.JSON(common.NewSuccessResponse(fiber.StatusCreated, "created", data))
}
