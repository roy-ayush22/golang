package main

import (
	"fmt"
	greetings "greeting"
)

func main() {
	message := greetings.Hello("Roman")
	fmt.Println(message)
}
