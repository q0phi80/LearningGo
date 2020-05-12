package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"checkNum:123,"amount":200,"category":["gift","clothing"]}`) // JSON that we don't know the structure
	var v interface{}
	json.Unmarshal(jsonData, &v) // Even though we don't know the JSON structure, we can unmarshal it into an interface
	fmt.Println(v)
}
