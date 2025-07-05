package api

import (
	"amf/models"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// var smfBaseURL string = "http://localhost:8081"
var smfBaseURL = os.Getenv("SMF_BASE_URL")

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
	//verify by TLS
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		fmt.Println("Failed to read cert.pem ")
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create reusable HTTP client
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90,
	}
	client := &http.Client{Transport: transport}
	response, err := client.Do(req)
	return response, err
}
