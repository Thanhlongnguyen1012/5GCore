package handler

import (
	"amf/api"
	"amf/models"

	"github.com/gin-gonic/gin"
)

func HandleN1N2Tranfer(c *gin.Context) {
	var request models.N1N2MessageTransferReqData
	err := c.ShouldBindJSON(&request)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	// 	return
	// }
	status, response := api.PostN1N2Tranfer(request, err)
	c.JSON(status, response)
}
