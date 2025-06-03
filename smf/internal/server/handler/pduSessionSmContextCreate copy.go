package handler

import (
	"fmt"
	"log"
	"net/http"
	"smf/internal/client"
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

	// Tạo Job với response channel
	responseChan := make(chan JobResult)
	job := Job{
		Payload:      Payload(request),
		ResponseChan: responseChan,
	}
	JobQueue <- job

	// Chờ kết quả từ worker
	result := <-responseChan

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if err == nil && result.Status == 200 {
		//defer client.SendPFCPEstablismentrequest()
		created := models.SMContextCreatedData{
			PduSessionID: request.PduSessionId,
			SNssai:       request.SNssai,
		}

		c.JSON(http.StatusCreated, created)

		//if result.Status == http.StatusCreated {
		// Run PFCP request in goroutine
		go func() {
			//time.Sleep(1 * time.Second)
			client.SendPFCPEstablismentrequest()
			log.Println("PFCP Session established successfully")
			data := models.N1N2MessageTransferReqData{
				PduSessionId: request.PduSessionId,
				SNssai:       request.SNssai,
				Dnn:          request.Dnn,
			}
			//time.Sleep(1 * time.Second)
			client.SendN1N2tranfer(data)
			log.Println("send N1N2Tranfer ok !")
		}()
	} else {
		fmt.Println("Failed to create SM Context")
	}
}
