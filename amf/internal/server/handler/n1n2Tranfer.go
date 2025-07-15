package handler

import (
	"amf/api"
	"amf/models"

	"github.com/gin-gonic/gin"
)

//var N1n2Count uint64

func HandleN1N2Tranfer(c *gin.Context) {
	var request models.N1N2MessageTransferReqData
	//atomic.AddUint64(&N1n2Count, 1)
	err := c.ShouldBindJSON(&request)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	// 	return
	// }
	status, response := api.PostN1N2Tranfer(request, err)
	c.JSON(status, response)
}
