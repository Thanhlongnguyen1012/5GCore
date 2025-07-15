package server

import (
	"fmt"
	"net/http"
	"smf/internal/server/handler"
	"smf/internal/server/resource"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	// Start JobQueue, Dispatcher for worker pool
	handler.JobQueue = make(chan handler.Job, handler.MaxQueue)
	dispatcher := handler.NewDispatcher(handler.MaxWorker)
	dispatcher.Run()
	// api.JobQueue = make(chan api.Job, api.MaxQueue)
	// dispatcher2 := api.NewDispatcher(api.MaxWorker)
	// dispatcher2.Run()
	//log.Println("Worker pool dispatcher started")
	// Start Prometheus metrics
	//prometheus.MustRegister(metric.HttpRequestsTotal)
	// Start server
	//s.r = gin.Default()
	s.r = gin.New()
	//s.r.Use(metric.PrometheusMiddleware())
	//resource.RouteSmContextCreate(s.r)
	/*s.r.Use(func(c *gin.Context) {
		fmt.Println("Protocol:", c.Request.Proto) // <-- In ra HTTP/1.1 hoặc HTTP/2.0
		c.Next()
	})
	*/
	resource.RouteSmContextCreate(s.r)
	h2s := &http2.Server{
		MaxConcurrentStreams: 10,
		IdleTimeout:          30 * time.Second,
	}
	server := &http.Server{
		Addr:    url,
		Handler: h2c.NewHandler(s.r, h2s),
	}
	err := server.ListenAndServe()
	//err := s.r.RunTLS(url, "cert.pem", "key.pem")
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
