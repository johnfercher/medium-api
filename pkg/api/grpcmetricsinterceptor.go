package api

import (
	"context"
	"time"

	"github.com/johnfercher/medium-api/pkg/observability/metrics/endpointmetrics"
	googleGrpc "google.golang.org/grpc"
)

type Interceptor struct{}

func NewInterceptor() *Interceptor {
	return &Interceptor{}
}

func (i *Interceptor) Intercept(ctx context.Context, req any, info *googleGrpc.UnaryServerInfo,
	handler googleGrpc.UnaryHandler,
) (any, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	latency := time.Since(start).Seconds()
	i.metrify(info, err, latency)

	return resp, err
}

func (i *Interceptor) metrify(info *googleGrpc.UnaryServerInfo, err error, latencyInMs float64) {
	metrics := endpointmetrics.Metrics{
		Protocol: "GRPC",
		Latency:  latencyInMs,
		Endpoint: info.FullMethod,
	}

	if err != nil {
		metrics.Failed = true
	} else {
		metrics.Failed = false
	}

	endpointmetrics.Send(metrics)
}
