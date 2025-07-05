package server

import (
	"fmt"
	"log"
	"smf/internal/server/resource"

	"smf/internal/server/handler"

	metric "smf/internal/server/metrics"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	// Start JobQueue, Dispatcher for worker pool
	handler.JobQueue = make(chan handler.Job, handler.MaxQueue)
	dispatcher := handler.NewDispatcher(handler.MaxWorker)
	dispatcher.Run()
	log.Println("Worker pool dispatcher started")
	// Start Prometheus metrics
	prometheus.MustRegister(metric.HttpRequestsTotal)
	// Start server
	s.r = gin.Default()
	s.r.Use(metric.PrometheusMiddleware())
	resource.RouteSmContextCreate(s.r)
	err := s.r.RunTLS(url, "cert.pem", "key.pem")
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
