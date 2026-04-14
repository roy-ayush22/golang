package main

import "fmt"

func main() {
	var name string
	var age int
	
	fmt.Print("enter your name: ")
	fmt.Scan(&name)
	
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	
	fmt.Println("hello", name, "you are", age, "years old!")
}