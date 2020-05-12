package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
A  program that can analyze an unknown JSON structure and,
for each field in the structure, print the data type and the JSON key-value pair.
*/
func main() {
	jsonData := []byte(`
	{
		"id": 2,  
    	"lname": "Asamoah",  
    	"fname": "Akwasi",  
    	"IsEnrolled": true,  
    	"grades":[100,76,93,50],  
    	"class":  
    	{  
      		"coursename": "World Lit",  
      		"coursenum": 101,  
      		"coursehours": 3  
    	}  
	}`)
	// Check whether the jsonData is valid JSON.
	if !json.Valid(jsonData) {
		fmt.Printf("[!] JSON is not valid: %s", jsonData)
		os.Exit(1)
	}
	// Declare an empty interface variable
	var v interface{}
	// Unmarshal jsonData into an empty interface
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Perform type switching on each value in the map
	data := v.(map[string]interface{})
	for k, v := range data {
		switch value := v.(type) {
		case string:
			fmt.Println("(string):", k, value)
		case float64:
			fmt.Println("(float64):", k, value)
		case bool:
			fmt.Println("(bool):", k, value)
		case []interface{}:
			fmt.Println("(slice):", k)
			for i, j := range value {
				fmt.Println(" ", i, j)
			}
		default:
			fmt.Println("(Unknown):", k, value)
		}
	}
}
