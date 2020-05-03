package main

import (
	"fmt"
)

func main() {
	songTitle := "Against_Foɔ"
	// Create a range loop that loops over the string, then capture the index and rune in variables
	for index, runeVal := range songTitle {
		fmt.Println(index, string(runeVal))
	}
}
