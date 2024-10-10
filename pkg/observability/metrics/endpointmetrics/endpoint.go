package endpointmetrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/johnfercher/medium-api/pkg/observability/log"

	"github.com/johnfercher/medium-api/pkg/observability/metrics/countermetrics"
	"github.com/johnfercher/medium-api/pkg/observability/metrics/histogrammetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// Labels.
	protocol            string = "protocol"
	endpoint            string = "endpoint"
	verb                string = "verb"
	pattern             string = "pattern"
	failed              string = "failed"
	error               string = "error"
	responseCode        string = "response_code"
	isAvailabilityError        = "is_availability_error"
	isReliabilityError         = "is_reliability_error"

	// Names.
	endpointRequestCounter string = "endpoint_request_counter"
	endpointRequestLatency string = "endpoint_request_latency"
)

var Helps = map[string]string{}

type Metrics struct {
	// Metric
	Protocol string
	Latency  float64

	// Labels
	Endpoint             string
	Verb                 string
	Pattern              string
	ResponseCode         int
	Failed               bool
	Error                string
	HasAvailabilityError bool
	HasReliabilityError  bool
}

func Send(metrics Metrics) {
	labels := map[string]string{
		protocol:            metrics.Protocol,
		endpoint:            metrics.Endpoint,
		verb:                metrics.Verb,
		pattern:             metrics.Pattern,
		responseCode:        fmt.Sprintf("%d", metrics.ResponseCode),
		failed:              fmt.Sprintf("%v", metrics.Failed),
		error:               metrics.Error,
		isAvailabilityError: fmt.Sprintf("%v", metrics.HasAvailabilityError),
		isReliabilityError:  fmt.Sprintf("%v", metrics.HasReliabilityError),
	}

	countermetrics.Increment(countermetrics.Metric{
		Name:   endpointRequestCounter,
		Labels: labels,
	})

	histogrammetrics.Observe(histogrammetrics.Metric{
		Name:  endpointRequestLatency,
		Value: metrics.Latency,
		Labels: map[string]string{
			endpoint: metrics.Endpoint,
		},
	})
}

func Start(ctx context.Context) {
	log.Info(ctx, "starting prometheus")
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		_ = http.ListenAndServe(":2112", nil)
	}()

	log.Info(ctx, "started prometheus")
}
