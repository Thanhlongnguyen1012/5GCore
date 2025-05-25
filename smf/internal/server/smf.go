package server

import (
	"fmt"
	"smf/internal/server/resource"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	s.r = gin.Default()
	resource.RouteSmContextCreate(s.r)
	err := s.r.Run(url)
	if err != nil {
		fmt.Println("smf start error")
	}
	fmt.Println("SMF start ok ")
}
