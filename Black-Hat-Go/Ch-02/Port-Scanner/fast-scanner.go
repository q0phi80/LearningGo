package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup // this acts as a synchronized counter
	for i := 1; i < 65535; i++ {
		wg.Add(1) // we increment the synchronized counter by 1 each time a goroutine is created to scan a port
		// Define goroutine for concurrency
		go func(j int) {
			defer wg.Done() // a defer to decrement the counter whenever one unit of the work has been performed
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			log.Printf("[+] Port %d open\n", j)
		}(i)
	}
	wg.Wait() // block until all the work has been done and the counter has returned to zero
}
