package countermetrics

import (
	"github.com/johnfercher/medium-api/pkg/observability/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Name   string
	Help   string
	Labels map[string]string
}

var createdMetrics = make(map[string]*prometheus.CounterVec)

func Increment(metric Metric) {
	go func() {
		labelsKey := metrics.GetLabelsKey(metric.Labels)

		opts := prometheus.CounterOpts{
			Name: metric.Name,
			Help: metric.Help,
		}

		if createdMetrics[metric.Name] == nil {
			counter := promauto.NewCounterVec(opts, labelsKey)
			createdMetrics[metric.Name] = counter
		}

		counter := createdMetrics[metric.Name]
		counter.With(metric.Labels).Inc()
	}()
}
