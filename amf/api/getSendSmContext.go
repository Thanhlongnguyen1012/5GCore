package api

import (
	"amf/models"
	"fmt"
	"io"
	"net/http"
)

func GetSmContextCreate(request models.SMContextCreateData) (int, interface{}) {
	resp, err := PostSmCreate(request)
	if err == nil && resp.StatusCode == 201 {
		created := "send sm context ok"
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
