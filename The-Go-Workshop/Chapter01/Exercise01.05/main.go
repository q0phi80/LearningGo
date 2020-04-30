package main

import (
    "fmt"
    "time"
)
func main() {
    // Declare each variable using the short variable declaration notation.
    Debug       := false
    LogLevel    := "info"
    startUpTime := time.Now()
    // Print the variables to the console
    fmt.Println(Debug, LogLevel, startUpTime)
}
