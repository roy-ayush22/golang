package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println(len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	b = [...]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	b = [...]int{100, 3:400, 500}
	fmt.Println((b))
}