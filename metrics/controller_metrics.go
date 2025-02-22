package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	BookOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "book_operations_total",
			Help: "Total number of book operations",
		},
		[]string{"operation"},
	)

	StatsOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stats_operations_total",
			Help: "Total number of stats operations",
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(BookOperationsTotal)
	prometheus.MustRegister(StatsOperationsTotal)
} 