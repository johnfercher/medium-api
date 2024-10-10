package chaos

import (
	"net/http"
	"time"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/johnfercher/medium-api/pkg/api"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type chaosHTTPHandler struct {
	inner             api.HTTPHandler
	baseSleepMs       float64
	latencyTendency   bool
	derivativeLatency float64
}

func NewChaosHTTPHandler(inner api.HTTPHandler, baseSleepMs float64) *chaosHTTPHandler {
	return &chaosHTTPHandler{
		inner:             inner,
		baseSleepMs:       baseSleepMs,
		latencyTendency:   BuildLatencyTendency(),
		derivativeLatency: 0,
	}
}

func (c *chaosHTTPHandler) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()
	c.sleep(c.baseSleepMs)

	err := c.generateError()
	if err != nil {
		log.Error(ctx, "injected error", field.Error(err))
		return nil, err
	}

	return c.inner.Handle(r)
}

func (c *chaosHTTPHandler) Name() string {
	return c.inner.Name()
}

func (c *chaosHTTPHandler) Pattern() string {
	return c.inner.Pattern()
}

func (c *chaosHTTPHandler) Verb() string {
	return c.inner.Verb()
}

// nolint:gomnd // magic number
func (c *chaosHTTPHandler) generateError() apierror.APIError {
	randValue := RandomFloat64(0, 100)
	if randValue < 10 {
		return c.getAvailabilityError()
	}

	if randValue < 25 {
		return c.getReliabilityError()
	}

	return nil
}

// nolint:gomnd // magic number
func (c *chaosHTTPHandler) getAvailabilityError() apierror.APIError {
	randValue := RandomFloat64(0, 100)
	if randValue < 33 {
		return apierror.New("service_unavailable", http.StatusServiceUnavailable)
	}

	if randValue < 66 {
		return apierror.New("internal_error", http.StatusInternalServerError)
	}

	return apierror.New("bad_gateway", http.StatusBadGateway)
}

// nolint:gomnd // magic number
func (c *chaosHTTPHandler) getReliabilityError() apierror.APIError {
	randValue := RandomFloat64(0, 100)
	if randValue < 33 {
		return apierror.New("bad_request", http.StatusBadRequest)
	}

	if randValue < 66 {
		return apierror.New("not_found", http.StatusNotFound)
	}

	return apierror.New("conflict", http.StatusConflict)
}

// nolint:gomnd // magic number
func (c *chaosHTTPHandler) sleep(ms float64) (appliedSleep float64) {
	positiveDerivative := c.getDerivativeWithTendency()
	jitter := GenerateJitter(ms, jitterPercent)

	if positiveDerivative {
		c.derivativeLatency = (c.derivativeLatency + jitter) / 2.0
	} else {
		c.derivativeLatency = (c.derivativeLatency - jitter) / 2.0
	}

	msDerivative := ms + c.derivativeLatency

	time.Sleep(time.Millisecond * time.Duration(msDerivative))
	return msDerivative
}

func (c *chaosHTTPHandler) getDerivativeWithTendency() bool {
	derivative := RandomBool()
	if c.latencyTendency != derivative {
		return RandomBool()
	}

	return derivative
}
