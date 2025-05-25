package api

import (
	"fmt"
	"io"
	"net/http"
	"smf/models"
)

func PostSmContextCreate(request models.SMContextCreateData) (int, interface{}) {
	resp, err := GetSessionManagementSubscription(request)
	if err == nil && resp.StatusCode == 200 {
		//defer client.SendPFCPEstablismentrequest()
		created := models.SMContextCreatedData{
			PduSessionID: request.PduSessionId,
			SNssai:       request.SNssai,
		}
		return http.StatusCreated, created
	}
	fmt.Println("Failed to create SM Context")
	var cause string
	if resp != nil && resp.Body != nil {
		body, _ := io.ReadAll(resp.Body)
		cause = string(body)
	} else {
		cause = "Unknown failure"
	}

	problem := models.ProblemDetails{
		Title:  "Failed Query SM Context at UDM",
		Status: http.StatusBadRequest,
		Cause:  cause,
	}
	return http.StatusInternalServerError, problem
}
