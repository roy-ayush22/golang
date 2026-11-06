package main

import (
	"fmt"
	greetings "greeting"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ayush")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
