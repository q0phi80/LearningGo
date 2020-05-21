package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{}

// Read function reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Println("Akwaaba...")
	fmt.Print("in>> ")
	return os.Stdin.Read(b)
}

// FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}

// Write function writes data to Stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Println("Medaase...")
	fmt.Print("out>> ")
	return os.Stdout.Write(b)

}
func main() {
	// Instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)
	// Use Go's copy function for copying input into an output
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("[!] Unable to read/write data")
	}
}
