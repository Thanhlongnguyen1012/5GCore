package api

import (
	"amf/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var smfBaseURL string = "http://localhost:8081"

func PostSmCreate(data models.SMContextCreateData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal err %v", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/nsmf-pdusession/v1/sm-contexts/", smfBaseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request creation error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	return response, err
}
