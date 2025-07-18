package handler

import (
	"net/http"
	"smf/models"

	"github.com/gin-gonic/gin"
)

var request models.SMContextCreateData

func HandlePDUSessionSmContextCreate(c *gin.Context) {
	//metrics for post createSmContext
	//metric.SMCreateRequestsTotal.Inc()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// Create Job with response channel
	responseChan := make(chan JobResult)
	job := Job{
		Payload:      Payload(request),
		ResponseChan: responseChan,
	}
	JobQueue <- job

	//  Waiting for result from worker
	result := <-responseChan

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.Status == 200 {
		//defer client.SendPFCPEstablismentrequest()
		created := models.SMContextCreatedData{
			PduSessionID: request.PduSessionId,
			SNssai:       request.SNssai,
		}

		c.JSON(http.StatusCreated, created)
		//} else {
		//fmt.Println("Failed to create SM Context")

		//if result.Status == http.StatusCreated {
		// Run PFCP request in goroutine
		// err := client.SendPFCPEstablismentrequest()
		// if err == nil {
		// 	job := api.Job{
		// 		Payload: api.Payload(models.N1N2MessageTransferReqData{
		// 			PduSessionId: request.PduSessionId,
		// 			SNssai:       request.SNssai,
		// 			Dnn:          request.Dnn,
		// 		}),
		// 		ResponseChan: make(chan api.JobResult),
		// 	}
		// 	api.JobQueue <- job
		// } else {
		// 	fmt.Println("pfcp Session establishment failed")
		// }
		// } else {
		// 	fmt.Println("Failed to create SM Context")
		// }
	}
}
