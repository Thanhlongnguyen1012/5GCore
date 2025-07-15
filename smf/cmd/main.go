package main

import (
	"fmt"
	"smf/internal/server"
)

func main() {
	//var wg sync.WaitGroup
	//SmfURL := "127.0.0.1:8081"
	//smfURL := os.Getenv("SMF_URL")
	fmt.Println("Start smf")
	//Start server
	var s server.Server
	s.Start(":8081")
	/*wg.Add(1)
	go func()
		defer wg.Done()
		s.Start(":8081")
	}()
	wg.Wait()
	*/
}
