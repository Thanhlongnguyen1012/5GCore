package main

import (
	"fmt"
	"smf/internal/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//SmfURL := "127.0.0.1:8081"
	//smfURL := os.Getenv("SMF_URL")
	fmt.Println("Start smf")
	//Start server
	var s server.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start("0.0.0.0:8081")
	}()
	wg.Wait()
}
