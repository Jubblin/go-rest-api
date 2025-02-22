package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ObjectBoxOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "objectbox_operations_total",
			Help: "Total number of ObjectBox operations",
		},
		[]string{"operation", "entity"},
	)

	ObjectBoxOperationDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "objectbox_operation_duration_seconds",
			Help:    "Duration of ObjectBox operations in seconds",
			Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1},
		},
		[]string{"operation", "entity"},
	)

	ObjectBoxEntityCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "objectbox_entity_count",
			Help: "Current number of entities in ObjectBox",
		},
		[]string{"entity"},
	)
)

func init() {
	prometheus.MustRegister(ObjectBoxOperationsTotal)
	prometheus.MustRegister(ObjectBoxOperationDuration)
	prometheus.MustRegister(ObjectBoxEntityCount)
} 