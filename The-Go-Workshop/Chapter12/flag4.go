package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Handling CTRL-C interruption
*/
func main() {
	sigs := make(chan os.Signal, 1)     // Create a channel of the os.Signal type. The Notify method works by sending values of the os.Signal type to a channel.
	done := make(chan bool)             // This channel is used to let us know when the program can exit.
	signal.Notify(sigs, syscall.SIGINT) // Receive the notification on the sigs channel.
	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL-C")
				fmt.Println("Some clean up is occuring.")
				done <- true // We send true to the done channel to indicate that we received the signal. This will stop the channel from blocking.
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught.")
	<-done // The <-done will be blocking until the program receives the signal.
	fmt.Println("Out of here.")
}
