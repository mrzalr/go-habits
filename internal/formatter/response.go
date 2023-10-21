package formatter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/common"
	"github.com/mrzalr/go-habits/internal/habit"
)

func SendSuccessResponse(c *fiber.Ctx, status common.Status, data interface{}) error {
	c.Status(status.Code)
	return c.JSON(common.NewSuccessResponse(status.Code, status.Message, data))
}

func SendErrorResponse(c *fiber.Ctx, err error, traceID string) error {
	status := errorMapper(err)

	c.Status(status.Code)
	return c.JSON(common.NewErrorResponse(status.Code, status.Message, traceID))
}

func errorMapper(err error) common.Status {
	errMap := map[error]common.Status{
		habit.ErrDataNotFound:   common.StatusNotFound,
		habit.ErrAlreadyStarted: common.StatusInternalServerError,
	}

	status, ok := errMap[err]
	if !ok {
		return common.StatusInternalServerError
	}

	status.Message = err.Error()
	return status
}
