package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nebulagate_requests_total",
			Help: "Total requests received",
		},
	)
	RequestsFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nebulagate_failed_requests_total",
			Help: "Total failed requests",
		},
	)
	BackendRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nebulagate_backend_requests_total",
			Help: "Requests per backend",
		},

		[]string{"backend"},
	)
)

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(
		RequestsTotal,
		RequestsFailed,
		BackendRequests,
	)
}
