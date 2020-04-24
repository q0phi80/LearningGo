package main

import (
  "fmt"
  "os"
)

func main() {
    f, err := os.Create("output.txt") // We create a file, assign a file pointer (f) and an error.
    if err != nil {
        panic("[!] Unable to create a file.")
    }
    // Close the file when the current scope is existed.
    defer f.Close()
    // Go has two functions for writing to files,a Write function (takes a byte buffer to write bytes to a file) and WriteString (takes strings and write to a file)
    cnt, err := f.WriteString("Hello, World!\n") // Write to the file, return the bytes written and an error.
    if err != nil {
        panic("[!] Unable to write to file.")
    }
    fmt.Printf("Wrote %d bytes\n", cnt) // Print the number of bytes written.
}
