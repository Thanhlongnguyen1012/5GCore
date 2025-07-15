package main

import (
	"amf/api"
	"amf/internal/client"
	"amf/internal/server"
	"amf/models"
	"log"
	"sync"
	"sync/atomic"
	"time"
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
	// request/s
	//const sendRate = 2000
	//duration ticker
	var durationTicker = 100 * time.Millisecond
	//var tickNum int = int(time.Second) / int(durationTicker)
	done := time.After(60 * time.Second)
	ticker := time.NewTicker(durationTicker)
	defer ticker.Stop()
	//amfBaseURL := "127.0.0.1:8080"
	//amfBaseURL := os.Getenv("AMF_BASE_URL")
	//start server
	go func() {
		var s server.Server
		s.Start(":8080")
	}()
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					//defer wg.Done()
					defer func() {
						if r := recover(); r != nil {
							log.Println("panic in goroutine:", r)
						}
						wg.Done()
					}()
					client.SendSmContextCreate(data)
				}()
			}
		case <-done:
			log.Println("create channel done")
			time.Sleep(30 * time.Second)
			wg.Wait()
			log.Println("tổng số request gửi đi trong 1s:", atomic.LoadUint64(&api.SmContextCount))
			log.Println("tổng số request thành công trong 1s:", atomic.LoadUint64(&api.N1n2Count))
			return
			//select {}
		}
	}
}
