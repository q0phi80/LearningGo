package main

import (
	"fmt"
)

func main() {
	// Declare a variable which is a slice of "strings" and initialize with data
	names := []string{"Kofi", "Akwasi", "Adwoa", "Yaw", "Abena"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
