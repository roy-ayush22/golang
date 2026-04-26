package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	target  int
	attempt int
}

func main() {
	game := Game{
		target: rand.Intn(101),
	}

	var guess int
	fmt.Println("guess a number between 0 and 100.")

	for {
		fmt.Print("guess the number:")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("invalid input! enter a valid integer", err)
			continue
		}

		game.attempt++

		if guess < game.target {
			fmt.Println("low!")
		} else if guess > game.target {
			fmt.Println("high!")
		} else {
			fmt.Println("correct, the number was:", game.target)
			fmt.Println("attempts:", game.attempt)
			break
		}
	}
}
