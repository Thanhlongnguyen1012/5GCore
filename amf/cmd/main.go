package main

import (
	"amf/internal/client"
	"amf/internal/server"
	"amf/models"
	"sync"
)

var data = models.SMContextCreateData{
	Supi:         "imsi-452040989692072",
	Gpsi:         "msisdn-84989692072",
	PduSessionId: 5,
	Dnn:          "v-internet",
	SNssai: &models.Snssai{
		Sst: 1,
		Sd:  "000001",
	},
	ServingNfId: "2ab2b5a9-68e8-4ee6-b939-024c109b520c",
	AnType:      "3GPP_ACCESS",
}

func main() {
	var wg sync.WaitGroup
	//amfBaseURL := "127.0.0.1:8080"
	//amfBaseURL := os.Getenv("AMF_BASE_URL")
	var s server.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start("0.0.0.0:8080")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		client.SendSmContextCreate(data)
	}()
	wg.Wait()
}
