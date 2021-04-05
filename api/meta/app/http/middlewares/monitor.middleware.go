package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

type PrometheusMonitor struct {
	ServiceName string
	APIRequestsCounter *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	RequestSize *prometheus.HistogramVec
	ResponseSize *prometheus.HistogramVec
}

func NewPrometheusMonitor(namespace, serviceName string) *PrometheusMonitor {
	APIRequestsCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name: "http_request_total",
			Help: "A counter for requests to the wrapped handler",
		},
		[]string{"handler", "method", "code", "micro_name"},
	)

	RequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name: "http_request_duration_seconds",
			Help: "A histogram of latencies for requests",
		},
		[]string{"handler", "method", "code", "micro_name"},
	)

	RequestSize := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name: "http_request_size_bytes",
			Help: "A histogram of request sizes for requests",
		},
		[]string{"handler", "method", "code", "micro_name"},
	)

	ResponseSize := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name: "http_response_size_bytes",
			Help: "A histogram of response sizes for requests",
		},
		[]string{"handler", "method", "code", "micro_name"},
	)

	prometheus.MustRegister(APIRequestsCounter, RequestDuration, RequestSize, ResponseSize)

	return &PrometheusMonitor{
		ServiceName: serviceName,
		APIRequestsCounter: APIRequestsCounter,
		RequestDuration: RequestDuration,
		RequestSize: RequestSize,
		ResponseSize: ResponseSize,
	}
}

func (m *PrometheusMonitor) PromMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		relativePath := c.Request.URL.Path
		start := time.Now()
		reqSize := computeApproximateRequestSize(c.Request)
		c.Next()
		duration := time.Since(start)
		code := fmt.Sprintf("%d", c.Writer.Status())
		m.APIRequestsCounter.With(prometheus.Labels{"handler": relativePath, "method": c.Request.Method, "code": code, "micro_name": m.ServiceName}).Inc()
		m.RequestDuration.With(prometheus.Labels{"handler": relativePath, "method": c.Request.Method, "code": code, "micro_name": m.ServiceName}).Observe(duration.Seconds())
		m.RequestSize.With(prometheus.Labels{"handler": relativePath, "method": c.Request.Method, "code": code, "micro_name": m.ServiceName}).Observe(float64(reqSize))
		m.ResponseSize.With(prometheus.Labels{"handler": relativePath, "method": c.Request.Method, "code": code, "micro_name": m.ServiceName}).Observe(float64(c.Writer.Size()))
	}
}

func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}