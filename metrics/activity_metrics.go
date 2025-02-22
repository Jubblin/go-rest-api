package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ActivityOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "activity_operations_total",
			Help: "Total number of activity operations",
		},
		[]string{"operation", "grid", "device"},
	)

	ActivityCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "activity_count",
			Help: "Current number of activities",
		},
		[]string{"grid", "device"},
	)

	ActivityLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "activity_operation_duration_seconds",
			Help:    "Duration of activity operations in seconds",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5},
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(ActivityOperationsTotal)
	prometheus.MustRegister(ActivityCount)
	prometheus.MustRegister(ActivityLatency)
} 