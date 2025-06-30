package server

import (
	"fmt"
	"log"
	"smf/internal/server/resource"
	"strconv"
	"time"

	"smf/internal/server/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Biến metric Prometheus
var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Tổng số lượng HTTP requests theo method, path, status",
		},
		[]string{"method", "path", "status"},
	)
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		path := c.FullPath() // sử dụng FullPath() để tránh phân mảnh bởi ID động

		if path == "" {
			path = c.Request.URL.Path // fallback nếu route không khớp
		}

		HttpRequestsTotal.WithLabelValues(method, path, status).Inc()
		log.Printf("Request %s %s [%s] took %v\n", method, path, status, duration)
	}
}

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	// Khởi tạo JobQueue, Dispatcher cho worker pool
	handler.JobQueue = make(chan handler.Job, handler.MaxQueue)
	dispatcher := handler.NewDispatcher(handler.MaxWorker)
	dispatcher.Run()
	log.Println("Worker pool dispatcher started")
	// Khởi tạo Prometheus metrics
	prometheus.MustRegister(HttpRequestsTotal)
	//khởi tạo server
	s.r = gin.Default()
	s.r.Use(PrometheusMiddleware())
	resource.RouteSmContextCreate(s.r)
	err := s.r.Run(url)
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
