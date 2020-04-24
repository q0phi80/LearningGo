package main

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

// Build a struct for a JSON structure
type Response struct {
    URL       string            `json:"url"`
    Headers   map[string]string `json:"headers"`
}
func main() {
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatalln("[!] Unable to get a response.")
    }
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("[!] Unable to read body content.")
    }
    respContent := Response{}
    json.Unmarshal(content, &respContent)
    fmt.Printf("From: %s\n", respContent.URL)
    for k, v := range respContent.Headers {
        fmt.Printf("%s: %s\n", k, v)
    }
}
