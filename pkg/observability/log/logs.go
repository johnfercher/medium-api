package log

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/johnfercher/medium-api/pkg/observability/log/field"
	"github.com/sirupsen/logrus"
)

var singletonLogger *logrus.Logger = nil

func NewLogger(level logrus.Level) *logrus.Logger {
	if singletonLogger != nil {
		return singletonLogger
	}

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(level)
	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))

	singletonLogger = logger

	return singletonLogger
}

func Info(ctx context.Context, msg string, fields ...*field.Field) {
	logger := GetLogger(ctx)
	logFields := buildFields(ctx, fields...)
	logger.WithFields(logFields).Info(msg)
}

func Error(ctx context.Context, msg string, fields ...*field.Field) {
	logger := GetLogger(ctx)
	logFields := buildFields(ctx, fields...)
	logger.WithFields(logFields).Error(msg)
}

func Warn(ctx context.Context, msg string, fields ...*field.Field) {
	logger := GetLogger(ctx)
	logFields := buildFields(ctx, fields...)
	logger.WithFields(logFields).Warning(msg)
}

func Debug(ctx context.Context, msg string, fields ...*field.Field) {
	logger := GetLogger(ctx)
	logFields := buildFields(ctx, fields...)
	logger.WithFields(logFields).Debug(msg)
}

func buildFields(ctx context.Context, fields ...*field.Field) logrus.Fields {
	logFields := logrus.Fields{}

	logFields["request_id"] = ctx.Value("request_id")

	for _, field := range fields {
		logFields[field.Key] = field.Value
	}

	return logFields
}
