package log

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey struct{}

func AddContext(ctx context.Context, logger *logrus.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func GetLogger(ctx context.Context) *logrus.Logger {
	logger := ctx.Value(contextKey{})
	return logger.(*logrus.Logger)
}

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, _ := uuid.NewRandom()

		ctx = AddContext(ctx, NewLogger(logrus.InfoLevel))
		ctx = context.WithValue(ctx, "request_id", id.String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
