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
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
)

var (
	SmContextCount uint64
	client         *http.Client
	smfBaseURL     = os.Getenv("SMF_BASE_URL")
	initOnce       sync.Once
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
func InitHttpClient() {
	client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http2.Transport{
			AllowHTTP: true, // Cho phép HTTP/2 không dùng TLS (h2c)
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				// Sử dụng kết nối TCP thường (không mã hóa)
				return net.Dial(network, addr)
			},
		},
	}
}
func PostSmCreate(data models.SMContextCreateData) (*http.Response, error) {
	atomic.AddUint64(&SmContextCount, 1)
	initOnce.Do(InitHttpClient)
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
	/*caCert, err := os.ReadFile("cert.pem")
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
	client = &http.Client{Transport: transport}
	*/
	response, err := client.Do(req)
	return response, err
}
