package handler

import (
	"fmt"
	"net/http"
	"smf/api"
	"smf/internal/client"
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
	if result.Error == nil && result.Status == 200 {
		//defer client.SendPFCPEstablismentrequest()
		created := models.SMContextCreatedData{
			PduSessionID: request.PduSessionId,
			SNssai:       request.SNssai,
		}

		c.JSON(http.StatusCreated, created)

		//if result.Status == http.StatusCreated {
		// Run PFCP request in goroutine
		err := client.SendPFCPEstablismentrequest()
		if err == nil {
			//log.Println("PFCP Session established successfully")
			/*data := models.N1N2MessageTransferReqData{
				PduSessionId: request.PduSessionId,
				SNssai:       request.SNssai,
				Dnn:          request.Dnn,
			}
			//time.Sleep(1 * time.Second)
			client.SendN1N2tranfer(data)
			*/
			//log.Println("send N1N2Tranfer ok !")
			job := api.Job{
				Payload: api.Payload(models.N1N2MessageTransferReqData{
					PduSessionId: request.PduSessionId,
					SNssai:       request.SNssai,
					Dnn:          request.Dnn,
				}),
				ResponseChan: make(chan api.JobResult),
			}
			api.JobQueue <- job
		} else {
			fmt.Println("pfcp Session establishment failed")
		}
	} else {
		fmt.Println("Failed to create SM Context")
	}
}
