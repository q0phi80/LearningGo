package main

import (
  "io/ioutil"
)

func main() {
    // ioutil allows to use one single function for both writing and reading.
    err := ioutil.WriteFile("output.txt", []byte("Hello, World!\n"), 0644) // We have to specify the file permission the file will have once its created.
    if err != nil {
        panic("[!] Unable to write to file.")
    }
}
