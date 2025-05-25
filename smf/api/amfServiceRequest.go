package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"smf/models"
)

var amfBaseURL string = "http://localhost:8080"

func PostN1N2Tranfer(request models.N1N2MessageTransferReqData) (*http.Response, error) {
	jsonData, err := json.Marshal(&request)
	if err != nil {
		fmt.Println("marshal JSON error")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/namf-comm/v1/ue-contexts/imsi-452040916843227/n1-n2-messages", amfBaseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("send post error")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
