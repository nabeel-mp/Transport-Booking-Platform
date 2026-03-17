package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricsRegisterOnce sync.Once
	httpRequestTotal    = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Total Request Count",
		},
		[]string{"method", "path", "status"},
	)
)

var httpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP Request latency",
	},
	[]string{"method", "path", "status"},
)

func PrometheusHTTP() fiber.Handler {
	return func(c fiber.Ctx) error {

		if c.Path() == "/metrics" {
			return c.Next()
		}

		metricsRegisterOnce.Do(func() {
			prometheus.MustRegister(httpRequestTotal, httpRequestDuration)
		})

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		method := c.Method()

		route := "unmateched"
		if r := c.Route(); r != nil {
			route = r.Path
		}

		status := strconv.Itoa(c.Response().StatusCode())

		httpRequestTotal.WithLabelValues(method, route, status).Inc()

		httpRequestDuration.WithLabelValues(method, route, status).Observe(duration.Seconds())

		return err
	}
}
