package main

import (
	"fmt"
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
	// Create buffer to hole input/output
	input := make([]byte, 4096) // Create a slice

	// Use reader to read the input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("[!] Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s) // Return the length of the data

	// Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("[!] Unable to writer dat")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
