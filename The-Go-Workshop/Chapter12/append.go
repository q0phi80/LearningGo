package main

import (
	"os"
)

/*
How to append data to a file.
*/
func main() {
	f, err := os.OpenFile("junk101.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write([]byte("Adding more stuff\n")); err != nil {
		panic(err)
	}
}
