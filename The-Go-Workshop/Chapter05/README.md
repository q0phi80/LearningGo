Function
Differenct parts are:
- func: Function Declaration
- Identifier: AKA the function name. You can have function without an identifier, this is called Anonymous function.
- Parameter list: Parameters are input values to a function. A parameter is a data that is required by the function to help solve the task of the function. Parameters are defined as: name, type (e.g. name string, age int). Parameters are optional for a function.
- Return types: List of data types such as Boolean, string, map or another function that can be returned.
- Function body: The coding statements between curly braces {}
- Function signature: The input parameters combined with the return types.
To execute a function, we must call it.
Functions parameters are local to the function, thus, they are only available to that function.
Functions typically accept inputs, perform some action on those inputs and then return result of those inputs.
Go allows the ability to ignore a variable being returned, called blank identifier
- e.g. _, err := file.Read(bytes)...in this example, the return value is ignored by using an underscore (_)
When there is extra data being returned from a function that does not provide any information that is needed by our program, such as the reading of a file, it is a good candidate for ignoring the return.
Good practice to name your return values...e.g. func greeting() (name string, age int) {}...name and age are the named return variables.
A variadic function is a function that accepts a variable number of argument values...func f(parameterName ...Type) The three dots (...) in front of the type is called pack operator.
The pack operator tells Go to store all the arguments of Type into parameterName.

Anonymous Functions
A function that doesn't have a function name.
Anonymous functions can do basically whatever a normal function in Go does, including accepting arguments and returning values.
How Anonymous Function looks like:
func main() {
  func() { // The anonymous function declaration
    fmt.Println("Greeting") // The Anonymous function body
  }() // Execution parantheses, which invoke the anonymous function.
}

To be able to pass arguments to an anonymous function, they must be supplied in the execution parentheses.
func main() {
  message := "Greeting"
  func(str string) { // Anonymous function with an input parameter of a string
    fmt.Println(str)
  }(message) // The argument messages being passed to the execution parantheses
}

An anonymous function can be saved to a variable.
func main() {
  f := func() {
    fmt.Println("Executing an anonymous function using a variable")
  }
  fmt.Println("Line after anonymous function")
  f() // Must include () after the variable to execute the function.
}

Defer
Deferred functions are commonly used for performing "clean-up" activities such as closing files, releasing resources, closing database connection and removal of configuration\temp files created by a program.
Defer statement can be utilized with anonymous functions.