package main

import (
  "fmt"
  "os"
)

func main() {
    r, err := os.Open("output.txt")
    if err != nil {
        panic("[!] Unable to open the file")
    }
    defer r.Close()
    // The Read function uses a byte buffer
    buf := make([]byte, 64) // We create the byte buffer and give it 64 bytes of size because don't have much data to read.
    cnt, err := r.Read(buf) // We read into a buffer.
    if err != nil {
        panic("[!] Unable to read the file.")
    }
    fmt.Printf("Read %d bytes\n", cnt)
    //fmt.Println(buf) // We print out the byte buffer.
    //fmt.Println(buf[:cnt])
    fmt.Println(string(buf[:cnt]))
}
