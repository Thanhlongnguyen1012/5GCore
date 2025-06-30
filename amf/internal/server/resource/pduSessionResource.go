package resource

import (
	"amf/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func RouteN1N2Tranfer(r *gin.Engine) {
	api := r.Group("/namf-comm")
	api.POST("/v1/ue-contexts/imsi-452040989692072/n1-n2-messages", handler.HandleN1N2Tranfer)
}
