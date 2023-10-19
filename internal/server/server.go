package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mrzalr/go-habits/pkg/configuration"
)

type server struct {
	app *fiber.App
	db  *sqlx.DB
	cfg *configuration.Configuration
}

func (s *server) Run() error {
	s.RegisterHandler()

	port := s.cfg.Http.Port
	if port == "" {
		return ErrInvalidPort
	}

	log.Printf("Listening on port %s\n", port)
	return s.app.Listen(fmt.Sprintf(":%s", port))
}

func New(db *sqlx.DB, cfg *configuration.Configuration) *server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return &server{
		app: app,
		db:  db,
		cfg: cfg,
	}
}
