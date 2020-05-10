package main

import (
	"errors"
	"fmt"
)

// Declare error variables using errors.New().
var (
	ErrHourlyRate  = errors.New("[!] Invalid hourly rate.")
	ErrHoursWorked = errors.New("[!] Invalid hours worked per week.")
)

func main() {
	pay, err := payDay(81, 50) // We call the payDay()
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 5) // We call the payDay()
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 50) // We call the payDay()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}
func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 { // Check whether the hourly rate is less 10 or greater 75
		return 0, ErrHourlyRate // If hourlyRate meets these conditions, return 0 and the custom error ErrHourlyRate
	}
	if hoursWorked < 0 || hoursWorked > 80 { // Check whehter hoursWorked is less than 0 or greater than 80.
		return 0, ErrHoursWorked // We return 0 and custom error if hours worked is less than 0 or greater than 80.
	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	}
	return hoursWorked * hourlyRate, nil // If the hourlyRate doesn't meet the conditions for checking 10 and 75, then we return the results of the hoursWorked multiplied by the hourlyRate and nil. We return nil for the error because there is no error.

}
