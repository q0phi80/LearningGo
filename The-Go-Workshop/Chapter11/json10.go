package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"mname,omitempty"` // Omitted if a value is not set
	FirstName     string   `json:"fname"`
	IsMarried     bool     `json:"-"`                  // Should not appear in the JSON
	IsEnrolled    bool     `json:"enrolled,omitempty"` // Omit if not set
	Courses       []course `json:"classes"`
}
type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

// This function will return a student struct
func newStudent(studentID int, lastName, middleInitial, firstName string, isMarried, isEnrolled bool) student {
	s := student{StudentId: studentID, LastName: lastName, MiddleInitial: middleInitial, FirstName: firstName, IsMarried: isMarried, IsEnrolled: isEnrolled}
	return s
}
func main() {
	s := newStudent(1, "Peprah", "S", "Ntim", false, false) // Use the newStudent() function to create a student struct and assign the result of the function to a variable s.
	student1, err := json.MarshalIndent(s, "", "    ")      // marshal s to JSON
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
	fmt.Println()
	s2 := newStudent(2, "Kwabena", "", "Antwi", true, true)
	c := course{Name: "World Lit", Number: 101, Hours: 3}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Biology", Number: 201, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Intro to Go", Number: 101, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))

}
