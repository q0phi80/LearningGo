package main

import (
	"fmt"
	"html/template" // changing from text/template to html/template will ensure that the user output will be escaped by this library.
	"net/http"
)

var content = `<html> 
<head> 
<title>Akwaaba</title> 
</head> 
<body> 
  <h1>Me ma wo Akwaaba</h1> 
  <h2> Kyerɛ w'adwen </h2> 
  {{.Comment}} 
  <form action="/" method="post"> 
    Ka w'adwen:<input type="text"name="input"> 
    <input type="submit" value="Submit"> 
  </form> 
</body> 
</html>`

// The struct will be used to wrap a user comment
type input struct {
	Comment string
}

// A handler to return the response of an HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	var userInput = &input{
		Comment: r.FormValue("input"),
	}
	t := template.Must(template.New("test").Parse(content))
	err := t.Execute(w, userInput)
	if err != nil {
		fmt.Println(err)
	}
}

// Run the HTTP server
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
