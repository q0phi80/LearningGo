package main

import (
  "fmt"
  "os"
  "log"
  "flag"
)

var name string

func main() {
    prefix := fmt.Sprintf("%s: ", os.Args[0])
    info, err := os.Create("info.log")
    if err != nil {
        log.Fatalln("[!] Failed to create a log file.")
    }
    infoLog := log.New(info, prefix, log.LstdFlags)
    errorLog := log.New(os.Stderr, prefix, log.LstdFlags)

    if name == "" {
        errorLog.Println("[!] No name supplied.")
        flag.Usage()
        os.Exit(1)
    }
    infoLog.Println("Program started!")
    infoLog.Println("Hello, ", name)
    infoLog.Println ("Program finished!")

}

func init() {
    flag.StringVar(&name, "name", "", "The name to say hello to.")
    flag.Parse()
}
