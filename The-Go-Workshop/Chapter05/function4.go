package main

import (
	"fmt"
	"strings"
)

/*
Mapping Index Values to Column Headers
A function to take a slice of columns headers from a CSV file
and print out a map of an index value of the headers we are interested in.
*/
func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	csvHdrCol(hdr2)
}
func csvHdrCol(header []string) {
	csvHeadersToColumIndex := make(map[int]string) // Assign a variable to a key-value pair of int and string...key(int) will be the index of header(string) column
	for i, v := range header {                     // Range over the header to process each string that is in the slice
		v = strings.TrimSpace(v) // Remove any trailing spaces in front of and after the string
		// Lower all the casing for exact matches.
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumIndex[i] = v
		case "hours worked":
			csvHeadersToColumIndex[i] = v
		case "hourly rate":
			csvHeadersToColumIndex[i] = v
		}

	}
	fmt.Println(csvHeadersToColumIndex)
}
