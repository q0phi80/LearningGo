package main

import (
	"fmt"
	"time"
)

/*
Create a function that tells the difference between the current time zone and the specified time zone.
The function will utilize the LoadLocation() function to specify the location based on which a variable will be set to a specific time.
The In() location will be used to convert a specific time value to a given time zone value
*/
func timeDiff(timezone string) (string, string) {
	Current := time.Now()
	RemoteZone, _ := time.LoadLocation(timezone)
	RemoteTime := Current.In(RemoteZone)
	fmt.Println("The current time is: ", Current.Format(time.ANSIC))
	fmt.Println("The timezone:", timezone, "time is:", RemoteTime)
	return Current.Format(time.ANSIC), RemoteTime.Format(time.ANSIC)
}
func main() {
	fmt.Println(timeDiff("America/Chicago"))
}
