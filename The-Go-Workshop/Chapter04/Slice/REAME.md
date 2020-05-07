SLICE
The size of a slice is its length, which we can find by usiing the len built-in function.
Go's built-in cap is used to determine the capacity of a slice.
Go's built-in append is used to add a new value to a slice.
You can't compare slices to one another but can compare slice to nil.

MAKE
Go's built-in make function allows you to set the length and capacity of a slice when creating it.
Syntax: make([]<sliceType>), <length>, <capacity>
When creating a slice using make, the capacity is optional but length is required.