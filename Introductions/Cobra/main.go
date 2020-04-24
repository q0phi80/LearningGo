// This is simple tool to demonstrate Cobra CLI for basic auth against a REST API.
package main

import (
  "fmt"
  "github.com/spf13/cobra"
  "io/ioutil"
  "log"
  "net/http"
)

// Create a command struct
// Create a command and assign a cobra struct.
var cmd = &cobra.Command{
    Use: "cobra", //Provide the name of our command
    Short: "This tool gets a URL with basic auth", // Provide a short description of the tool.
    Run: func (cmd *cobra.Command, args []string)  {// A function which provides the Cobra command and also the commandline arguments.
        //fmt.Println("Hello!") // We provide a print statement to make the function works.
        // Make URL a requirement for the user to provide.
        if len(args) != 1 {
            log.Fatalln("[!] Must provide a URL.")
        }
        // After a URL is validated, we create a client for the connection by creating a client struct
        client := http.Client{}
        // Create a web request, this case, GET request
        req, err := http.NewRequest("GET", args[0], nil) // pass in the URL as an argument.
        if err != nil {
            log.Fatalln("[!] Unable to make a GET request.")
        }
        // We check if the user provided a username and password before we request the basic auth.
        if username != "" && password !="" {
            req.SetBasicAuth(username, password)
        }
        resp, err := client.Do(req)
        if err != nil {
            log.Fatalln("[!] Unable to get the response.")
        }
        defer resp.Body.Close() // We want to make sure that once we exist the scope, the body is closed.
        // We want to get the content of the web response body.
        content, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln("[!] Unable to read body content.")
        }
        // We then print the body content. Since the content is a byte size, we need to cast it to string.
        fmt.Println(string(content))
    },
}

// We need to add flags so that a user can pass a username and password.
var username, password string

func main () {
      cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for the credentials.")
      cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for the credentials.")
      cmd.Execute() // Execute our command.

}
