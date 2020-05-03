package main

import (
	"fmt"
)

func main() {
	username := "fluffy_wabbit"
	runes := []rune(username)
	// Loop and print each byte of string
	for i := 0; i < len(username); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
