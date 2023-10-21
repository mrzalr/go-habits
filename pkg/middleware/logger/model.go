package logger

import (
	"time"

	"go.uber.org/zap"
)

type loggerModel struct {
	app           string
	version       string
	method        string
	status        int
	header        string
	uri           string
	body          string
	response      string
	traceID       string
	detailedError string
	responseTime  time.Duration
}

func (m *loggerModel) GenerateLogFields() []zap.Field {
	fields := []zap.Field{
		zap.String("app", m.app),
		zap.String("version", m.version),
		zap.String("method", m.method),
		zap.Int("status", m.status),
		zap.String("header", m.header),
		zap.String("uri", m.uri),
		zap.String("body", m.body),
		zap.String("response", m.response),
		zap.Duration("response_time", m.responseTime),
		zap.String("trace_id", m.traceID),
	}

	if len(m.detailedError) != 0 {
		fields = append(fields, zap.String("detailed_error", m.detailedError))
	}

	return fields
}
