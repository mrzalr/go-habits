package cmd

import (
	"log"

	"github.com/mrzalr/go-habits/internal/server"
	"github.com/mrzalr/go-habits/pkg/configuration"
	"github.com/mrzalr/go-habits/pkg/database/mysql"
)

func StartApplication() {
	cfg := &configuration.Configuration{}
	err := cfg.SetConfig("config")
	if err != nil {
		log.Fatalf("log fatal set config : %v", err)
	}

	db, err := mysql.New(cfg)
	if err != nil {
		log.Fatalf("log fatal connect to mysql : %v", err)
	}

	s := server.New(db, cfg)
	log.Fatal(s.Run())
}
