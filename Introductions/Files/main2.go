package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/spf13/viper"
)

func main() {
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        log.Fatalln("[!] Unable to get a response.")
    }
    defer resp.Body.Close()
    viper.SetConfigType("json")
    if err := viper.ReadConfig(resp.Body); err != nil {
        log.Fatalln("[!] Unable to ready body content.")
    }
    //fmt.Println(viper.AllSettings())
    m := viper.GetStringMapString("headers")
    for k, v := range m {
        fmt.Printf("%s: %s\n", k, v)
    }
}
