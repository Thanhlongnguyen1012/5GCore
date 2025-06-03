package server

import (
	"fmt"
	"log"
	"smf/internal/server/resource"

	"smf/internal/server/handler"

	"github.com/gin-gonic/gin"
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
	//khởi tạo server
	s.r = gin.Default()
	resource.RouteSmContextCreate(s.r)
	err := s.r.Run(url)
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
