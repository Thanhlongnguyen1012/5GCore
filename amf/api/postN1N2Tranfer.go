package api

import (
	"amf/models"
	"net/http"
)

func PostN1N2Tranfer(request models.N1N2MessageTransferReqData, err error) (int, interface{}) {
	if err == nil {
		return http.StatusOK, nil
	}
	var cause string
	cause = "Invalid JSON"
	problem := models.ProblemDetails{
		Title:  "error Invalid JSON",
		Status: http.StatusBadRequest,
		Cause:  cause,
	}
	return http.StatusInternalServerError, problem
}
