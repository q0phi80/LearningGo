package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) { // Accept two channels
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0 // If the port is closed, send zero
			continue
		}
		conn.Close()
		results <- p // If the port is open, send the port to the results
	}
}
func main() {
	ports := make(chan int, 100)
	results := make(chan int) // Create a separate channel to communicate the results from the worker
	var openports []int       // Use slice to store the results for later sorting
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port) // If the port doesn't equal 0, its appended to the slice
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)             // After closing the channels, use sort to sort the slice of open ports
	for _, port := range openports { // Loop over the slice and print the open ports to screen
		fmt.Printf("Port %d is open\n", port)
	}
}
