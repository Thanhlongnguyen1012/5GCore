package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"smf/api"
	"smf/models"
)

var (
	MaxWorker = 128
	MaxQueue  = 128
)

// var UdmBaseURL string = "http://localhost:8082"
var (
	UdmBaseURL = os.Getenv("UDM_BASE_URL")
	amfBaseURL = os.Getenv("AMF_BASE_URL")
)

type Payload models.SMContextCreateData

func (p *Payload) GetSessionManagementSubscription() (*http.Response, error) {
	api.InitOnce.Do(api.InitHttpClient)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", UdmBaseURL, p.Supi), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
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
	// client := &http.Client{Transport: transport}
	response, err := api.Client.Do(req)
	return response, err
}
func (p *Payload) SendN1N2tranfer(data models.N1N2MessageTransferReqData) (*http.Response, error) {
	api.InitOnce.Do(api.InitHttpClient)
	jsonData, err := json.Marshal(&data)
	if err != nil {
		fmt.Println("marshal JSON error")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/namf-comm/v1/ue-contexts/imsi-452040989692072/n1-n2-messages", amfBaseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("send post error")
	}

	resp, err := api.Client.Do(req)
	//metrics for n1n2
	//metric.N1N2RequestsTotal.Inc()
	if err != nil {
		fmt.Println("http request n1n2 failed ")
		return nil, err
	}
	return resp, err
}

// Job represents the job to be run
type Job struct {
	Payload      Payload
	ResponseChan chan JobResult
}
type JobResult struct {
	Status   int
	Response any
	Error    error
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				resp, err := job.Payload.GetSessionManagementSubscription()
				if err != nil {
					job.ResponseChan <- JobResult{
						Status: http.StatusInternalServerError,
						Error:  err,
					}
					continue
				}
				defer resp.Body.Close()

				var result any
				body, err := io.ReadAll(resp.Body)
				if err == nil {
					json.Unmarshal(body, &result)
				}

				job.ResponseChan <- JobResult{
					Status:   resp.StatusCode,
					Response: result,
				}

			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
