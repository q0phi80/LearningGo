package main

import (
  "fmt"
  "time"
)
func main() {
  for {
    fmt.Println("Tick")
    time.Sleep(1 * time.Second)
    break // Adding a break tells the for-loop to stop after it prints the first Tick message.
  }
}
