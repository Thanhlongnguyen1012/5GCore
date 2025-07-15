package server

import (
	"amf/internal/server/resource"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	r *gin.Engine
}

func (s Server) Start(url string) {
	//s.r = gin.Default()
	s.r = gin.New()
	/*s.r.Use(func(c *gin.Context) {
		fmt.Println("Protocol:", c.Request.Proto) // <-- In ra HTTP/1.1 hoặc HTTP/2.0
		c.Next()
	})*/
	resource.RouteN1N2Tranfer(s.r)
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    url,
		Handler: h2c.NewHandler(s.r, h2s),
	}
	err := server.ListenAndServe()
	//err := s.r.RunTLS(url, "cert.pem", "key.pem")
	if err != nil {
		fmt.Println("AMF start error")
	}
	fmt.Println("AMF start ok ")
}
