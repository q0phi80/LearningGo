package main

import (
	"fmt"
	"math"
)

/*
Numeric type conversion
*/
// A function that returns a string
func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14
	m := fmt.Sprintf("int8 = %v > in64 	= %v\n", i8, int64(i8))      // Convert from a small int type to a larger int type.
	m += fmt.Sprintf("int  = %v > in8     	= %v\n", i, int8(i))      // Convert from an int that is 1 above int8's maximum size. This should cause an overflow to int8's minimum size.
	m += fmt.Sprintf("int8 = %v > float32 	= %v\n", i8, float64(i8)) // Convert int8 to float64.
	m += fmt.Sprintf("f64  = %v > int 	= %v\n", f64, int(f64))       // Convert a float to an int
	return m                                                         // Return the message
}
func main() {
	fmt.Print(convert())
}
