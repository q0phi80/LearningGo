package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Create a function that allows you to calculate the duration of the execution between two time.Time variables
and return a string that tells you how long the execution took to complete.
*/
func elapsedTime(start time.Time, end time.Time) string {
	Elapsed := end.Sub(start)
	Hours := strconv.Itoa(int(Elapsed.Hours()))
	Minutes := strconv.Itoa(int(Elapsed.Minutes()))
	Seconds := strconv.Itoa(int(Elapsed.Seconds()))
	return "The total execution time elapsed is: " + Hours + " hour(s) and " + Minutes + " minutes(s) and " + Seconds + " second(s)!"
}
func main() {
	Start := time.Now()
	time.Sleep(2 * time.Second)
	End := time.Now()
	fmt.Println(elapsedTime(Start, End))
}
