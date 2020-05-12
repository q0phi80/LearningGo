package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	name := "q0phi80"
	log.Println("Demo app")
	log.Printf("%s is here!", name)
	log.Print("Run")
}
