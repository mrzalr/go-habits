package common

import "github.com/gofiber/fiber/v2"

type Status struct {
	Code    int
	Message string
}

var (
	StatusOk                  = Status{fiber.StatusOK, "success"}
	StatusCreated             = Status{fiber.StatusCreated, "created"}
	StatusNotFound            = Status{fiber.StatusNotFound, "not found"}
	StatusInternalServerError = Status{fiber.StatusInternalServerError, "internal server error"}
)
