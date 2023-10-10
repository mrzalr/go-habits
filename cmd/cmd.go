package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	welcomeHandler "github.com/mrzalr/go-habits/internal/welcome/handler"
	welcomeRepo "github.com/mrzalr/go-habits/internal/welcome/repository/mysql"
	welcomeUcase "github.com/mrzalr/go-habits/internal/welcome/usecase"
	"github.com/mrzalr/go-habits/pkg/configuration"
	"github.com/mrzalr/go-habits/pkg/database/mysql"
)

func StartApplication() {
	config := configuration.Configuration{}
	err := config.SetConfig("config")
	if err != nil {
		log.Fatalf("log fatal set config : %v", err)
	}

	db, err := mysql.New(config)
	if err != nil {
		log.Fatalf("log fatal connect to mysql : %v", err)
	}

	app := fiber.New()

	welcomeRepository := welcomeRepo.New(db)
	welcomeUsecase := welcomeUcase.New(welcomeRepository)
	welcomeHandler.New(app, welcomeUsecase)

	port := config.Http.Port
	log.Printf("Listening on port %d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
