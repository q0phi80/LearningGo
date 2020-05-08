package main

type Employee struct {
	Id                  int
	FirstName, LastName string
}
type developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}
