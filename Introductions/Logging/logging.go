package main

import (
  "log"
  "flag"
)

var name string

func main() {
    
    if name == "" {
        flag.Usage()
        log.Fatalln("[!] No name supplied.")
    }
    log.Println("Program started!")
    log.Println("Hello, %s\n", name)
    log.Println ("Program finished!")

}

func init() {
    flag.StringVar(&name, "name", "", "The name to say hello to.")
    flag.Parse()
}
