package main

import (
	"fmt"
	"smf/internal/server"
	"sync"
)

/*var data = models.N1N2MessageTransferReqData{
	PduSessionId: 99,
}*/

func main() {
	var wg sync.WaitGroup
	//SmfURL := "127.0.0.1:8081"
	//smfURL := os.Getenv("SMF_URL")
	fmt.Println("Start smf")
	//Khởi chạy server
	var s server.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start("0.0.0.0:8081")
	}()
	/*wg.Add(1)
	go func() {
		defer wg.Done()
		client.SendN1N2tranfer(data)
	}()
	*/
	wg.Wait()
}
