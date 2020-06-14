package personal

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Register custom prometheus metrics for entity routes
var (
	indexCounters = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "index",
			Help: "Number of requests for index page.",
		},
		[]string{"index"},
	)
)

// Register high cardinality route variables to be replaced by strings
// this avoids congestion on the metrics endpoint that prometheus
// scrapes.
var (
	URLReplacers = map[string]string{
		"resume": "resume",
		"about":  "about",
	}
)

//nolint:gochecknoinits
func init() {
	// register counters for GET and POST requests with Prometheus's default registry
	prometheus.MustRegister(indexCounters)
}
