package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"smf/models"
)

var (
	MaxWorker = 10
	MaxQueue  = 20
)

// var UdmBaseURL string = "http://localhost:8082"
var UdmBaseURL = os.Getenv("UDM_BASE_URL")

type Payload models.SMContextCreateData

func (p *Payload) GetSessionManagementSubscription() (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", UdmBaseURL, p.Supi), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	return response, err
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
