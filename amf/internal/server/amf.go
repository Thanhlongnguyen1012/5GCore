package server

import (
	"amf/internal/server/resource"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	s.r = gin.Default()
	resource.RouteN1N2Tranfer(s.r)
	err := s.r.Run(url)
	if err != nil {
		fmt.Println("AMF start error")
	}
	fmt.Println("AMF start ok ")
}
