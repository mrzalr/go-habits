package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mrzalr/go-habits/pkg/configuration"
	"github.com/mrzalr/go-habits/pkg/logger"
)

func Log(cfg *configuration.Configuration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeStart := time.Now()
		c.Next()

		request := c.Request()
		response := c.Response()
		method := request.Header.Method()
		status := response.StatusCode()
		header := request.Header.Header()
		uri := request.URI().Path()
		body := request.Body()
		respBody := response.Body()

		traceID := string(response.Header.Peek("X-Trace-ID"))
		if len(traceID) == 0 {
			traceID = uuid.New().String()
		}

		_loggerModel := loggerModel{
			app:          cfg.App,
			version:      cfg.Version,
			method:       string(method),
			status:       status,
			header:       string(header),
			uri:          string(uri),
			body:         string(body),
			response:     string(respBody),
			traceID:      traceID,
			responseTime: time.Since(timeStart),
		}

		logger.Info(":", _loggerModel.GenerateLogFields()...)
		return nil
	}
}
