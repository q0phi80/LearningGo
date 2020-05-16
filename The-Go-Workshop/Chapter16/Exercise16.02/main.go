package main

import (
	"log"
	"sync" // Will be used for WaitGroup
)

func sum(from, to int, wg *sync.WaitGroup, res *int) { // We add a parameter wg with a pointer to the sync.WaitGroup along with the result parameter
	*res = 0 // Set the value of what is held by the res pointer to 0.
	for i := from; i <= to; i++ {
		*res += i
	}
	wg.Done() // We tell the WaitGroup that is Goroutine is completed and then we return
	return
}
func main() {
	s1 := 0
	wg := &sync.WaitGroup{} // A pointer to the WaitGroup is created
	wg.Add(1)               // Notify the WaitGroup that is one Goroutine running
	go sum(1, 100, wg, &s1) // Create new Goroutine calculating the sum. The sum() function will call .Done() method to notify the WaitGroup of its completion
	// Wait for the Goroutine to finish
	wg.Wait()
	log.Println(s1)
}
