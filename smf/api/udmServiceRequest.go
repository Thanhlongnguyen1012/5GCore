package api

import (
	"fmt"
	"net/http"
	"smf/models"
)

var UdmBaseURL string = "http://localhost:8082"

func GetSessionManagementSubscription(data models.SMContextCreateData) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", UdmBaseURL, data.Supi), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	return response, err
}
