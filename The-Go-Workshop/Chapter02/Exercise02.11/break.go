package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create an empty loop to loop forever if not stopped
	for {
		// Use Intn from the rand package to pick a random number btn 0 and 8
		r := rand.Intn(1000)
		// If the random number is divisible by 3, print "Skip" and skip the rest of the loop using continue
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		}
		// If the random number is divible by 2, pring "Stop" and stop the loop using break
		if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		// If the random number is neither of those conditions, print the number
		fmt.Println(r)
	}
}
