package main

import (
    "fmt"
    "time"
)
// A function that returns three values.
func getConfig() (bool, string, time.Time) {
    return false, "info", time.Now() // Return three initial values, separated by a comma.
}

func main() {
    Debug, LogLevel, startUpTime := getConfig() //Use variable declaration to capture the values returned from the getConfig three new variables.
    fmt.Println(Debug, LogLevel, startUpTime)
}
