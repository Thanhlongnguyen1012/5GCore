package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	metric "smf/internal/server/metrics"
	"smf/models"
	"strconv"
)

// var amfBaseURL string = "http://localhost:8080"
var amfBaseURL = os.Getenv("AMF_BASE_URL")

func PostN1N2Tranfer(request models.N1N2MessageTransferReqData) (*http.Response, error) {
	jsonData, err := json.Marshal(&request)
	if err != nil {
		fmt.Println("marshal JSON error")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/namf-comm/v1/ue-contexts/imsi-452040989692072/n1-n2-messages", amfBaseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("send post error")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	//tăng giá trị mỗi khi gửi n1n2
	metric.N1N2RequestsTotal.Inc()
	metric.HttpRequestsTotal.
		WithLabelValues(req.Method, req.URL.Path, strconv.Itoa(resp.StatusCode)).
		Inc()
	return resp, err
}
