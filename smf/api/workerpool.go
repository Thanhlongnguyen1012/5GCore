package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"smf/models"
)

var (
	MaxWorker = 128
	MaxQueue  = 128
)

// var UdmBaseURL string = "http://localhost:8082"
type Payload models.N1N2MessageTransferReqData

func (p *Payload) SendN1N2tranfer() (*http.Response, error) {
	InitOnce.Do(InitHttpClient)
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal JSON error")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/namf-comm/v1/ue-contexts/imsi-452040989692072/n1-n2-messages", amfBaseURL), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("send post error")
	}

	resp, err := Client.Do(req)
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
				resp, err := job.Payload.SendN1N2tranfer()
				if err != nil {
					job.ResponseChan <- JobResult{
						Status: http.StatusInternalServerError,
						Error:  err,
					}
					job.ResponseChan <- JobResult{
						Status: resp.StatusCode,
					}

					continue
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
