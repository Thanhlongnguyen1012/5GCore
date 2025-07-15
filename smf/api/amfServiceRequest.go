package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"smf/models"
	"sync"

	"golang.org/x/net/http2"
)

// var amfBaseURL string = "http://localhost:8080"
var (
	amfBaseURL = os.Getenv("AMF_BASE_URL")
	Client     *http.Client
	InitOnce   sync.Once
)

/*
	func InitHttpClient() {
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

		Client = &http.Client{
			Transport: tr,
		}
	}
*/
func InitHttpClient() {
	Client = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true, // Cho phép HTTP/2 không dùng TLS (h2c)
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				// Sử dụng kết nối TCP thường (không mã hóa)
				return net.Dial(network, addr)
			},
		},
	}
}

func PostN1N2Tranfer(request models.N1N2MessageTransferReqData) (*http.Response, error) {
	InitOnce.Do(InitHttpClient)
	jsonData, err := json.Marshal(&request)
	if err != nil {
		fmt.Println("marshal JSON error")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/namf-comm/v1/ue-contexts/imsi-452040989692072/n1-n2-messages", amfBaseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("send post error")
	}
	//verify by TLS
	// caCert, err := os.ReadFile("cert.pem")
	// if err != nil {
	// 	fmt.Println("Failed to read cert.pem ")
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// // Create reusable HTTP client
	// transport := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		RootCAs: caCertPool,
	// 	},
	// 	MaxIdleConns:        100,
	// 	MaxIdleConnsPerHost: 100,
	// 	IdleConnTimeout:     90,
	// }
	// client = &http.Client{Transport: transport}
	resp, err := Client.Do(req)
	//metrics for n1n2
	//metric.N1N2RequestsTotal.Inc()
	if err != nil {
		fmt.Println("http request failed ")
		return nil, err
	}
	//metric.HttpRequestsTotal.
	//	WithLabelValues(req.Method, req.URL.Path, strconv.Itoa(resp.StatusCode)).
	//	Inc()
	return resp, err
}
