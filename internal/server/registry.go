package server

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mrzalr/go-habits/internal/habit/handler"
	habitRepo "github.com/mrzalr/go-habits/internal/habit/repository/mysql"
	habitUcase "github.com/mrzalr/go-habits/internal/habit/usecase"
	"github.com/mrzalr/go-habits/pkg/middleware"
	"github.com/mrzalr/go-habits/pkg/middleware/logger"
)

func (s *server) setupMiddlewares() {
	s.app.Use(recover.New())
	s.app.Use(logger.Log(s.cfg))
	s.app.Use(middleware.ErrorHandler)
}

func (s *server) RegisterHandler() {
	s.setupMiddlewares()

	habitRepository := habitRepo.New(s.db)
	habitUsecase := habitUcase.New(habitRepository)
	handler.New(s.app, habitUsecase)
}
