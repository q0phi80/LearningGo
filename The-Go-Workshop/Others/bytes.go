package main

import (
	"fmt"
)

func main() {
	username := "fluffy_wabbit"
	// Loop and print each byte of string
	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
}
