package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a channel of the os.Signal type that will be used to receive the notifications from the Notify method
	sigs := make(chan os.Signal, 1)
	// Add a done channel to be used to let us know when the program can exit
	done := make(chan bool)
	// The Notify method works by sending values of the os.Signal type to a channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP)
	// Create an anonymous function as a goroutine
	go func() {
		// Create an infinite loop where we receive a value from the sigs channel and store it in the s variable
		for {
			s := <-sigs
			// Create a switch statement that evaluates what is received from the channel
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- true
			case syscall.SIGTSTP:
				fmt.Println()
				fmt.Println("Someone pressed CTRL-Z")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- true
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught(ctrl-z, ctrl-c)")
	<-done
	fmt.Println("Out of here")
}

// A function to mimic a clean up process
func cleanUp() {
	fmt.Println("Simulating clean up")
	for i := 0; i < 10; i++ {
		fmt.Println("Deleting Files.. Not Really.", i)
		time.Sleep(1 * time.Second)
	}
}
