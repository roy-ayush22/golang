package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(101)

	var guess int
	fmt.Println("guess a number between 0 and 100.")

	for {
		fmt.Print("guess the number:")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("invalid input! enter a valid integer", err)
			continue

		}

		if guess < num {
			fmt.Println("low!")
		} else if guess > num {
			fmt.Println("high!")
		} else {
			fmt.Println("correct! the number was:", num)
			break
		}
	}
}
