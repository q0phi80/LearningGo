package main

import(
  "fmt"
)
// We can return values from functions
func getHello(name string) string { // We have to specify the return value type, which in this case is string.
    return fmt.Sprintf("Hello, %s\n", name) // Instead of printing, we can return the string we want to print. We use Sprintf which will use formating to print the string and return that string so that we can call and print the string in the main function.
}

func main() {
  s := getHello("Asamoah") // We call the function and pass in a name.
  fmt.Println(s) // We print out the string that was returned from getHello.
}
