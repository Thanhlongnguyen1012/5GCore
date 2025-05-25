package resource

import (
	"smf/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func RouteSmContextCreate(r *gin.Engine) {
	r.POST("/nsmf-pdusession/v1/sm-contexts/", handler.HandlePDUSessionSmContextCreate)
}
