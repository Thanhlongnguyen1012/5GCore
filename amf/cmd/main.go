package main

import (
	"amf/internal/client"
	"amf/internal/server"
	"amf/models"
	"sync"
)

var data = models.SMContextCreateData{
	Supi:         "imsi-452040916843227",
	Gpsi:         "msisdn-84867220452",
	PduSessionId: 5,
	Dnn:          "v-internet",
	SNssai: &models.Snssai{
		Sst: 1,
		Sd:  "000001",
	},
	ServingNfId: "",
	AnType:      "3GPP_ACCESS",
}

func main() {
	var wg sync.WaitGroup
	amfBaseURL := "localhost:8080"
	var s server.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start(amfBaseURL)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		client.SendSmContextCreate(data)
	}()
	wg.Wait()
}
