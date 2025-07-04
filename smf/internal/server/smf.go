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
	// Khởi tạo JobQueue, Dispatcher cho worker pool
	handler.JobQueue = make(chan handler.Job, handler.MaxQueue)
	dispatcher := handler.NewDispatcher(handler.MaxWorker)
	dispatcher.Run()
	log.Println("Worker pool dispatcher started")
	// Khởi tạo Prometheus metrics
	prometheus.MustRegister(metric.HttpRequestsTotal)
	//khởi tạo server
	s.r = gin.Default()
	s.r.Use(metric.PrometheusMiddleware())
	resource.RouteSmContextCreate(s.r)
	err := s.r.Run(url)
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
