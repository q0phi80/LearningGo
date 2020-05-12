package main

import (
	"os"
)

func main() {
	f, err := os.Create("test.txt") // Create an empty file.
	if err != nil {
		panic(err) // Better to panic and then exit because the defer function will not run if we do an os.Exit(1) with a function that has a defer function.
	}
	defer f.Close()
}
