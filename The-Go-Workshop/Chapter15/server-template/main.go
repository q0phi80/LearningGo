package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tplStr = `
<html>
	<h1>Customer {{.ID}}</h1>
	{{if .ID}}
		<p>Details:</p>
		<ul>
			{{if .Name}}<li>Name: {{.Name}}</li>{{end}}
			{{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
			{{if .Age}}<li>Age: {{.Age}}</li>{{end}}
		</ul>
		{{else}}
		<p>Data not available</p>
	{{end}}
</html>
`

// The following will fill the template
// A struct with attributes needed by the template
type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

// A handler function with a variable to the map of values in the querystring
func Hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	cust := Customer{}
	// Grab the passed value from the URL
	id, ok := v1["id"]
	if ok {
		cust.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}
	name, ok := v1["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}
	surname, ok := v1["surname"]
	if ok {
		cust.Surname = strings.Join(surname, ",")
	}
	age, ok := v1["age"]
	if ok {
		cust.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}
	tmpl, _ := template.New("test").Parse(tplStr)
	tmpl.Execute(w, cust)
}
func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
