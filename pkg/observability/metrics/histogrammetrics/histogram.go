package histogrammetrics

import (
	"github.com/johnfercher/medium-api/pkg/observability/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Name   string
	Help   string
	Value  float64
	Labels map[string]string
}

var createdMetrics = make(map[string]*prometheus.HistogramVec)

func Observe(metric Metric) {
	go func() {
		labelsKey := metrics.GetLabelsKey(metric.Labels)

		opts := prometheus.HistogramOpts{
			Name:    metric.Name,
			Help:    metric.Help,
			Buckets: GetDefaultBucket(),
		}

		if createdMetrics[metric.Name] == nil {
			histogram := promauto.NewHistogramVec(opts, labelsKey)
			createdMetrics[metric.Name] = histogram
		}

		histogram := createdMetrics[metric.Name]
		histogram.With(metric.Labels).Observe(metric.Value)
	}()
}

// nolint:gomnd // magic number
func GetDefaultBucket() []float64 {
	return prometheus.LinearBuckets(0.05, 0.050, 20)
}
