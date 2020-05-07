package main

import (
	"fmt"
)

func getDays() []string {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var newDays []string
	newDays = append(days, "Test")
	for s := range newDays {
		// newDays = append(days, "Test")
		fmt.Println("New Week :", s)
	}
	return newDays
}
func main() {
	newDays := getDays()
	fmt.Println("New Week :", newDays)
}
