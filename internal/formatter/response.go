package formatter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrzalr/go-habits/internal/common"
	m_habit "github.com/mrzalr/go-habits/internal/habit/model"
	"github.com/mrzalr/go-habits/pkg/validator"
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
		// Common
		validator.ErrStructValidation: common.StatusBadRequest,

		// Habit
		m_habit.ErrDataNotFound:   common.StatusNotFound,
		m_habit.ErrAlreadyStarted: common.StatusInternalServerError,
		m_habit.ErrBadRequest:     common.StatusBadRequest,
	}

	status, ok := errMap[err]
	if !ok {
		return common.StatusInternalServerError
	}

	status.Message = err.Error()
	return status
}
