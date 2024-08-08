package api

import (
	"net/http"
	"time"

	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
	"github.com/johnfercher/medium-api/pkg/encode"
	"github.com/johnfercher/medium-api/pkg/observability/metrics/endpointmetrics"
)

type HttpHandlerAdapter interface {
	AdaptHandler() func(w http.ResponseWriter, r *http.Request)
}

type metricsHandlerAdapter struct {
	handler HTTPHandler
}

func NewMetricsHandlerAdapter(handler HTTPHandler) *metricsHandlerAdapter {
	return &metricsHandlerAdapter{
		handler: handler,
	}
}

func (m *metricsHandlerAdapter) AdaptHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.execute(w, r)
	}
}

func (m *metricsHandlerAdapter) execute(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response, err := m.handler.Handle(r)

	latency := time.Since(start).Seconds()
	m.metrify(response, err, latency)

	if err != nil {
		http.Error(w, err.Name(), err.Code())
		return
	}

	encode.WriteJSONResponse(w, response.Object(), response.Code())
}

// nolint:gomnd // magic number
func (m *metricsHandlerAdapter) metrify(response apiresponse.APIResponse, err apierror.APIError, latencyInMs float64) {
	metrics := endpointmetrics.Metrics{
		Latency:  latencyInMs,
		Endpoint: m.handler.Name(),
		Verb:     m.handler.Verb(),
		Pattern:  m.handler.Pattern(),
	}

	if err != nil {
		metrics.Failed = true
		metrics.Error = err.Name()
		metrics.ResponseCode = err.Code()
		if err.Code() >= 500 {
			metrics.HasReliabilityError = false
			metrics.HasAvailabilityError = true
		} else {
			metrics.HasReliabilityError = true
			metrics.HasAvailabilityError = false
		}
	} else {
		metrics.Failed = false
		metrics.ResponseCode = response.Code()
	}

	endpointmetrics.Send(metrics)
}
