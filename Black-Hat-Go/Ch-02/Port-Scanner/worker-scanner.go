package main

import (
	"fmt"
	"sync"
)

/*
In this program, we will use a pool of goroutinges to manage the concurrent work being performed.
Using a for loop, we will create a certain number of worker goroutines as a resource pool and then use a channel to provide work, from witin the main() function.
A program to create 100 workers, consumes a channel of int and prints them on the screen.
*/
// Use WaitGroup function to block execution
// A worker function for processing work
// The channel will be used receive work and the WaitGroup will be used to track when a single work item has been completed.
func worker(ports chan int, wg *sync.WaitGroup) { // two arguments: a channel of type int and a pointer to a WaitGroup
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

// Main function manages the workload and provide work to the worker function
func main() {
	ports := make(chan int, 100) // Create a channel and a buffer of 100 for maintaining and tracking work for multiple producers and consumers
	var wg sync.WaitGroup
	// A loop to start the desired number of workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		ports <- i // Send the port on the ports channel to the worker.
	}
	wg.Wait()
	close(ports) // After all the work has been completed, close the channel.
}
