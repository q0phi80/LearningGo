package main

import (
	"fmt"
	"time"
)

/*
 Create a function called whatstheclock to demonstrate how you can create a function that wraps a nice,
 formatted time.Now() function and returns the date in an ANSIC format.
*/
func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}
func main() {
	fmt.Println("The script has started at: ", whatstheclock())
	fmt.Println("Saving the world...")
	time.Sleep(2 * time.Second)
	fmt.Println("The script has completed at: ", whatstheclock())
}
