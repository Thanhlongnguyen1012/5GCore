package handler

import (
	"log"
	"net/http"
	"smf/api"
	"smf/internal/client"
	"smf/models"
	"time"

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
	if fw, ok := c.Writer.(http.Flusher); ok {
		fw.Flush() // đẩy ngay buffer xuống client
	}
	if status == http.StatusCreated {
		// Run PFCP request in goroutine
		go func() {
			time.Sleep(2 * time.Second)
			client.SendPFCPEstablismentrequest()
			log.Println("PFCP Session established successfully")
			data := models.N1N2MessageTransferReqData{
				PduSessionId: 1, // <-- Thêm dấu phẩy nếu struct có nhiều trường
			}
			time.Sleep(1 * time.Second)
			client.SendN1N2tranfer(data)
			log.Println("send N1N2Tranfer ok !")
		}()
	}
}
