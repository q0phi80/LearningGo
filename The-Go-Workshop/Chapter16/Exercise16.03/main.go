package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	// A loop to add each number to the total
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i))
	}
	wg.Done()
	return
}
func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)
	// Four Goroutines to perform the sum over four ranges
	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)
	// The code that waits for the routines to complete and print the result
	wg.Wait()
	log.Println(s1)
}
