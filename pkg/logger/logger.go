package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.Logger
	loggerTDR *zap.Logger
)

func createFolderIfNotExists(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func buildLogger(cfg LoggerConfig, encodeCfg zapcore.EncoderConfig, dirPath, filePath string) (*zap.Logger, error) {
	ec := zapcore.NewJSONEncoder(encodeCfg)
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))

	if cfg.SaveLogFile {
		filePath = fmt.Sprintf("%s/%s.log", dirPath, filePath)
		err := createFolderIfNotExists(dirPath)
		if err != nil {
			return nil, err
		}

		logFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}

		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile),
		)
	}

	core := zapcore.NewCore(ec, writeSyncer, zap.InfoLevel)
	return zap.New(core), nil
}

func New(_loggerCfg LoggerConfig) error {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	sysFilePath := fmt.Sprintf("sys-%s", time.Now().Format("2006.01.02"))
	logger, err = buildLogger(_loggerCfg, cfg, _loggerCfg.SysLogFileLocation, sysFilePath)
	if err != nil {
		return err
	}

	tdrFilePath := fmt.Sprintf("tdr-%s", time.Now().Format("2006.01.02"))
	loggerTDR, err = buildLogger(_loggerCfg, cfg, _loggerCfg.TDRLogFileLocation, tdrFilePath)
	if err != nil {
		return err
	}

	return nil
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

func InfoTDR(message string, field ...zap.Field) {
	loggerWithField := loggerTDR.With(field...)
	loggerWithField.Info(message)
}
