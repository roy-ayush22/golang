package main

import "fmt"

func main()  {
	// print 10 numbers from 1
	i := 1
	for i < 11 {
		fmt.Println(i)
		i = i + 1
	}

	// same as above
	for j := 1; j < 11; j++ {
		fmt.Println(j)
	}

	
	for i := range 10 {
		fmt.Println(i)
	}

	// prints "loop" and breaks out of loop
	for {
		fmt.Println("loop")
		break
	}

	// print even numbers
	for m := range 10 {
		if m % 2 == 0 {
			fmt.Println(m)
		}
	}

	// print odd numbers
	for n := range 10 {
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	}	

}