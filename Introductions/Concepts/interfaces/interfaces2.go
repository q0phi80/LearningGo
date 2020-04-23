/*
Interfaces can also be used as abstractions for different types that have the same function
in their method sets.
*/
package main

import (
  "fmt"
)
// A pserson type
type BasePerson struct {
    First string
    Last  string
}
// An Employee type which is composed of BasePerson type
type Employee struct {
    BasePerson
    Salary  int
    Boss    *Manager
}
// A Manager type which composes of just the Employee type.
type Manager struct {
    Employee
}
// A person interface with only one function, GetName()
type Person interface {
    GetName() string
}
// We can add a function called SayHello and use it on any person, whether they are an employee or a manager.
func SayHello(p Person) {
    fmt.Printf("Hello, %s\n", p.GetName())
}
// Create a method for our Employee to get name.
func (e Employee) GetName() string {
    return e.First
}
// Create a method for our Manager to ge name.
func (m Manager) GetName() string {
    return m.First
}

func main() {
    m := &Manager { // We create a Manager that is an Employee, has a BasePerson, has a Salary and no Boss.
        Employee{
            BasePerson: BasePerson{
                First:  "Kofi",
                Last:   "Asamoah",
            },
            Salary: 400000,
            Boss: nil,
        },
    }
    e := &Employee{ // We create an Employee that has a BasePerson, a Salary and the Boss is the boss created above.
        BasePerson: BasePerson {
            First:  "Yaw",
            Last:   "Sarpong",
        },
        Salary: 30000,
        Boss:   m,
    }
    // We can call out the GetName function here and it will use the Person interface to use the same function.
    SayHello(m)
    SayHello(e)
}
