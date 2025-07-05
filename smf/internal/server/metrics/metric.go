package metric

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// metrics var
var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Tổng số lượng HTTP requests theo method, path, status",
		},
		[]string{"method", "path", "status"},
	)
)
var (
	// metrics SMCreateContext (AMF → SMF)
	SMCreateRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "smf_smcreate_requests_total",
			Help: "Tổng số lần AMF gửi SMCreateContext (POST /sm-context) đến SMF",
		},
	)

	// metrics N1N2 transfers (SMF → AMF)
	N1N2RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "smf_n1n2_requests_total",
			Help: "Tổng số lần SMF gửi N1N2MessageTransfer (POST /n1-n2-messages) đến AMF",
		},
	)
)

func init() {
	prometheus.MustRegister(SMCreateRequestsTotal, N1N2RequestsTotal)
}
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		path := c.FullPath()

		if path == "" {
			path = c.Request.URL.Path
		}

		HttpRequestsTotal.WithLabelValues(method, path, status).Inc()
		log.Printf("Request %s %s [%s] took %v\n", method, path, status, duration)
	}
}
