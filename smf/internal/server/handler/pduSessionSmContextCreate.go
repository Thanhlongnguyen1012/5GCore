package handler

import (
	"net/http"
	"smf/api"
	"smf/models"

	"github.com/gin-gonic/gin"
)

func HandlePDUSessionSmContextCreate(c *gin.Context) {
	var request models.SMContextCreateData
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	status, response := api.PostSmContextCreate(request)
	c.JSON(status, response)
}
