package main

import (
  "fmt"
  "flag"
  "os"
)

func main() {
    var name string
    flag.StringVar(&name, "name", "", "The name the user should enter.")
    flag.Parse()

    if name == "" {
        fmt.Println("[!] You must enter a username.")
        flag.Usage()
        os.Exit(1)
    }
    fmt.Printf("Hello, %s\n", name)
}
