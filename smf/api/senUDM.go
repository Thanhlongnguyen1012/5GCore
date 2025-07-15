package api

import (
	"fmt"
	"net/http"
	"os"
	"smf/models"
)

var (
	UdmBaseURL = os.Getenv("UDM_BASE_URL")
)

func GetSessionManagementSubscription(data models.SMContextCreateData) (*http.Response, error) {
	InitOnce.Do(InitHttpClient)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", UdmBaseURL, data.Supi), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := Client.Do(req)
	return response, err
}
