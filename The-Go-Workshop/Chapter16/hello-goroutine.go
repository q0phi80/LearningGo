package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Akwaaba!")

}

/*
In order to call the hello() function as a Goroutine, you can do this:
go hello()
*/
// Example
func main() {
	fmt.Println("Start")
	go hello()
	fmt.Println("End")
}

/*
The code starts by printing Start, then it calls the hello() function.
Then, the execution goes straight to printing End without waiting for the hello() function to complete.
No matter how long it takes to run the hello() function,
the main() function will not care about the hello() function as these functions will run independently.
*/
