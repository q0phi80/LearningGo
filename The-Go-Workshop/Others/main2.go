package main

import (
	"fmt"
)

func main() {
	songTitle := "Against_Foɔ"
	// Length of a string
	fmt.Println("Bytes:", len(songTitle))
	fmt.Println("Runes:", len([]rune(songTitle)))
	// Limit to 5 characters
	fmt.Println(string(songTitle[:5]))
	fmt.Println(string([]rune(songTitle)[:5]))
}
