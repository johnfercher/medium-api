package api

import (
	"context"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/sirupsen/logrus"
	googleGrpc "google.golang.org/grpc"
)

type LoggerInterceptor struct{}

func NewLoggerInterceptor() *LoggerInterceptor {
	return &LoggerInterceptor{}
}

func (i *LoggerInterceptor) Intercept(ctx context.Context, req any, _ *googleGrpc.UnaryServerInfo,
	handler googleGrpc.UnaryHandler,
) (any, error) {
	logger := log.NewLogger(logrus.InfoLevel)
	ctx = log.AddContext(ctx, logger)

	return handler(ctx, req)
}
