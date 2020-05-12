package main

import (
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("Using the Write function.\n")) // We take the string and convert it into a slice of bytes.
	f.WriteString("Using the WriteString function.\n")
}
