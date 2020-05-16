package main

import (
	"fmt"
	"sync"
)

/*
 Calculate a sum using several workers.
 A worker is essentially a function, and we
 will be organizing these workers into a single struct
*/
type Worker struct {
	in, out chan int
	sbw     int
	mtx     *sync.Mutex
}

// Create methods and increase the number of workers
func (w *Worker) readThem() {
	w.sbw++
	go func() {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial
		w.mtx.Lock() // Lock the routine
		w.sbw--      // Reduce the counter on the sub-workers safely
		if w.sbw == 0 {
			close(w.out) // Close the output channel
		}
		w.mtx.Unlock() // Unlock the execution to allow the program to carry on
	}()
}

// A function to return the sum
func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1) // Spawn only one routine
	go func() {
		for i := range w.out {
			total += i
		}
		wg.Done()
	}()
	wg.Wait() // Wait for the routine to finish and return the result
	return total
}

// Main function sets up the variables for the worker and its sub-workers
func main() {
	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	wrNum := 10
	out := make(chan int)
	wrk := Worker{in: in, out: out, mtx: mtx}
	for i := 1; i <= wrNum; i++ {
		wrk.readThem()
	}
	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)
	// Wait for the result and print it out
	res := wrk.gatherResult()
	fmt.Println(res)
}
