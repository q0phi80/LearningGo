package main

import (
  "fmt"
  "flag"
  "os"
  "time"
)
// We move the variable into Global scope
var debug bool
var name string
var wait time.Duration // How long Go should wait before it does what we want it to do.

func main() {
      if name == "" {
        fmt.Println("[!] You must enter a username.")
        flag.Usage()
        os.Exit(1)
    }

    if debug {
        fmt.Printf("Going to wait for %v\n", wait)
    }
    time.Sleep(wait) // Wait for the specified time before printing hello...
    fmt.Printf("Hello, %s\n", name)
}

// init function can be placed in any Go file.
// Any code within the init function is called prior to the code executing.
func init() { // We can take our commandline parsing code and put them within the init function.
    flag.BoolVar(&debug, "debug", false, "Turn on debugging output.")
    flag.StringVar(&name, "name", "", "The name the user should enter.")
    defaultWait, err := time.ParseDuration("5s")
    if err != nil {
        panic("[!] Could not parse default wait time.")
    }
    flag.DurationVar(&wait, "wait-time", defaultWait, "Time to wait before print.")
    flag.Parse()

}
