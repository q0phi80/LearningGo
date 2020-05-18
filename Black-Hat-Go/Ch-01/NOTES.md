## The go build Command
### Syntax: go build main.go OR go build -o hello main.go
- This command compiles the application, including any packages and their dependencies, without installing the results.
- This creates a binary file to be executed.
- To reduce the file size of the produced binary, you can add additional flags during build
  - *go -o hello -ldflags "-w -s" main.go*

## Cross-Compiling
- The build command allows for cross-compiling a go program for multiple Operating Systems and architectures.
- Reference: https://golang.org/doc/install/source#environment/
- You need to set a constraint to cross-compile a go program.
- The constraints include GOOS (the Operating System) and GOARCH (the architecture)
- Constraints can be introduced via command line, code comments of a file suffix naming convention
- Example cross-compile for Windows OS on a 64-architecture:
  - *GOOS="windows" GOARCH="amd64" go build -o akwaaba -ldflags "-w -s" maing.go*

## Slices
- Slices are like arrays that can be dynamically resized and passed to functions.
 - An example syntax to define, initialize and work with a slice **s**
 - *var s = make([]string, 0)*
 - *s = append(s, "some string")*
- The **make()** function is used to initialize a variable
- The **append()** function is used to add a new item to a slice 

## Maps
- Maps are associative arrays, with key/value pairs to be able to lookup values for a unique key
- An example syntax to define, initialize and work with a map **m**
 - *var m = make(map[string]string)*
 - *m["some key"] = "some value"...This code adds the key/value pair to the map **m**

## Pointers
- A **pointer** points to a particular area in memory and allows you to retrieve the value stored there
- You use the **&** operator to retrieve the address in memory of some variable and the * operator to deference the address
- An example syntax:
 - *var count = int(42)* - define an integer
 - *ptr := &count* - create a pointer...this returns the address of the **count** variable
 - *fmt.Println(*ptr)* - deference the variable and print the value of **count** to console
 - *ptr = 100 - use * operator to assign a new value to the memory location pointed to by ptr, which changes the value
 - fmt.Println(count)

## Structs
- A **struct** is a type used to define new data types by specifying the type's associated fields and methods
- Example struct for a **Person** type:
```
 type Person struct { - Define a new struct containing two fields: a string named Name and an int named Age
   Name string
   Age int
 }
 func (p *Person) SayHello() { - Define a method called SayHello on the Person type assigned to a variable p
   fmt.Println("Hello,", p.Name)
 }
 func main() {
   var guy = new(Person) - initialize a new Person
   guy.Name = "q0phi80" - assigns the name q0phi80 to the person
   guy.SayHello()
 }
```
- types and fields that begin with a capital letter are exported and accessible outside the package
- types and fields that begin with a lowercase letter are private and accessible only within the package

## for loop
- An example code to loop over a collection such as a slice
```
nums := []int{2,3,4,6,8} - initialize a slice of integers named num
for idx, val := range nums { - use range to iterate over the slice.
  fmt.Println(k, v)
}
```
 - The **range** keyword returns two values: the current index and a copy of the current value at that index
 - If the index is not need, **idx** in the loop can be replaced with an underscore **(_)** to tell Go we won't need it
