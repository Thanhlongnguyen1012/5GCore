package handler

import (
	"fmt"
	"net/http"
	"os"
	"smf/api"
	"smf/internal/client"
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

func (p *Payload) PduSessionEstablishment() (*http.Response, error) {
	resp, err := api.GetSessionManagementSubscription(request)
	if err != nil {
		fmt.Println("Failed to send udm")
		return nil, err
	}
	if resp.StatusCode == 200 && resp.Body != nil {
		err := client.SendPFCPEstablismentrequest()
		if err != nil {
			fmt.Println("Failed to send UPF")
			return nil, err
		}
		data := models.N1N2MessageTransferReqData{
			PduSessionId: request.PduSessionId,
			SNssai:       request.SNssai,
			Dnn:          request.Dnn,
		}
		err = client.SendN1N2tranfer(data)
		if err != nil {
			fmt.Println("Failed to send AMF")
		}
		fmt.Println("Send n1n2 ok !")
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
				resp, err := job.Payload.PduSessionEstablishment()
				if err != nil {
					job.ResponseChan <- JobResult{
						Status: http.StatusInternalServerError,
						Error:  err,
					}
					continue
				}

				job.ResponseChan <- JobResult{
					Status:   resp.StatusCode,
					Response: nil,
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
