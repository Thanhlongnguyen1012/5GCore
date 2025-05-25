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
	amfURL := "localhost:8081"
	fmt.Println("Start smf")
	//Khởi chạy server
	var s server.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start(amfURL)
	}()
	/*wg.Add(1)
	go func() {
		defer wg.Done()
		client.SendN1N2tranfer(data)
	}()
	*/
	wg.Wait()
}
