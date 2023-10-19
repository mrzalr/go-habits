package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	err    error
)

func init() {
	if logger != nil {
		return
	}

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableCaller = true

	logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
}

func Info(message string, field ...zap.Field) {
	loggerWithField := logger.With(field...)
	loggerWithField.Info(message)
}

func Warn(message string, field ...zap.Field) {
	loggerWithField := logger.With(field...)
	loggerWithField.Warn(message)
}

func Fatal(message string, field ...zap.Field) {
	loggerWithField := logger.With(field...)
	loggerWithField.Fatal(message)
}
