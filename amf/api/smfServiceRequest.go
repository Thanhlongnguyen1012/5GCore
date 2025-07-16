package api

import (
	"amf/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
)

var (
	SmContextCount uint64
	smfBaseURL     = os.Getenv("SMF_BASE_URL")
)

/*func initHttpClient() {
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		fmt.Println("Failed to read cert.pem:", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
		MaxIdleConnsPerHost: 1,
		IdleConnTimeout:     90,
		ForceAttemptHTTP2:   true, // Đảm bảo HTTP/2 nếu server hỗ trợ
		DisableKeepAlives:   false,
	}

	client = &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
}*/

// var smfBaseURL string = "http://localhost:8081"
var (
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
)

func PostSmCreate(data models.SMContextCreateData) (*http.Response, error) {
	atomic.AddUint64(&SmContextCount, 1)
	//initOnce.Do(InitHttpClient)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal err %v", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/nsmf-pdusession/v1/sm-contexts/", smfBaseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request creation error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	return response, err
}
