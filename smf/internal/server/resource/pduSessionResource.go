package resource

import (
	"smf/internal/server/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RouteSmContextCreate(r *gin.Engine) {
	r.POST("/nsmf-pdusession/v1/sm-contexts/", handler.HandlePDUSessionSmContextCreate)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
