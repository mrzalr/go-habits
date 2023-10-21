package cmd

import (
	"fmt"
	"log"

	"github.com/mrzalr/go-habits/internal/server"
	"github.com/mrzalr/go-habits/pkg/configuration"
	"github.com/mrzalr/go-habits/pkg/database/mysql"
	"github.com/mrzalr/go-habits/pkg/logger"
)

func StartApplication() {
	cfg := &configuration.Configuration{}
	err := cfg.SetConfig("config")
	if err != nil {
		log.Fatalf("CONFIG: unable to set config - %s", err.Error())
	}

	err = logger.New(logger.LoggerConfig{
		SysLogFileLocation: cfg.SysLogFileLocation,
		TDRLogFileLocation: cfg.TDRLogFileLocation,
		SaveLogFile:        cfg.SaveLogFile,
	})
	if err != nil {
		log.Fatalf("LOGGER: unable to create logger - %s", err.Error())
	}

	db, err := mysql.New(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("DB: failed to connect to database - %s", err.Error()))
	}

	s := server.New(db, cfg)
	err = s.Run()
	if err != nil {
		logger.Fatal(fmt.Sprintf("SERVER: unable to start the server - %s", err.Error()))
	}
}
