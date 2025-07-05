package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	metric "smf/internal/server/metrics"
	"smf/models"
	"strconv"
)

// var UdmBaseURL string = "http://localhost:8082"
var UdmBaseURL = os.Getenv("UDM_BASE_URL")

func GetSessionManagementSubscription(data models.SMContextCreateData) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", UdmBaseURL, data.Supi), nil)
	if err != nil {
		return nil, err
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
	// metrics
	metric.HttpRequestsTotal.WithLabelValues(req.Method, req.URL.Path, strconv.Itoa(response.StatusCode)).Inc()
	return response, err
}
