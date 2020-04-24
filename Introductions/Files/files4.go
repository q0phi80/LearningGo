package main

import (
  "io/ioutil"
  "fmt"
)

func main() {
    // ioutil allows to use one single function for both writing and reading.
    buf, err := ioutil.ReadFile("output.txt")
    if err != nil {
        panic("[!] Unable to read file.")
    }
    fmt.Println(string(buf))
}
