package controller

import (
	"time"

	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var latencyHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "request_latency_seconds",
	Help:    "Request latency in seconds.",
	Buckets: prometheus.DefBuckets,
})

var errorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "error_simulate_count",
	Help: "Simulate error count",
}, []string{"status_code"})

func init() {
	prometheus.MustRegister(latencyHistogram)
	prometheus.MustRegister(errorCounter)
}

func LatencyOutputs(c echo.Context) error {
	latency := rand.Intn(1000)
	time.Sleep(time.Duration(latency) * time.Millisecond)

	latencyHistogram.Observe(float64(latency) / 1000.0)

	return c.JSON(200, map[string]interface{}{
		"latency": latency,
	})
}

func ErrorSimulateCount(c echo.Context) error {
	statusCode := c.Param("statusCode")
	errorCounter.WithLabelValues(statusCode).Inc()

	switch statusCode {
	case "500":
		return c.JSON(500, map[string]interface{}{
			"error": "Simulated error count",
		})
	case "400":
		return c.JSON(400, map[string]interface{}{
			"error": "Simulated error count",
		})
	case "404":
		return c.JSON(404, map[string]interface{}{
			"error": "Simulated error count",
		})
	default:
		return c.JSON(500, map[string]interface{}{
			"error": "Simulated error count",
		})
	}
}
