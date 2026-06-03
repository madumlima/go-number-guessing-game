package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(101)

	var n int
	var difficultyLevel string
	var numberOfAttempts int

	fmt.Println("Welcome to the Number Guessing Game! ")
	fmt.Println("I'm thinking of a number between 1 and 100")

	fmt.Println("\nPlease select the difficulty level: ")
	fmt.Println("1 - Easy (10 chances)")
	fmt.Println("2 - Medium (5 chances)")
	fmt.Println("3 - Hard (3 chances)")

	fmt.Print("\nEnter your choice: ")
	fmt.Scanln(&n)

	var err error
	switch n {
	case 1:
		difficultyLevel = "Easy"
		numberOfAttempts = 10
	case 2:
		difficultyLevel = "Medium"
		numberOfAttempts = 5
	case 3:
		difficultyLevel = "Hard"
		numberOfAttempts = 3
	default:
		err = errors.New("invalid choice")

	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game\n", difficultyLevel)

	var guess int
	for i := 0; i < numberOfAttempts; i++ {
		fmt.Print("\nEnter your guess: ")
		fmt.Scanln(&guess)

		if guess == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", i+1)
			return
		}

		if guess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %d", guess)
			continue
		}

		fmt.Printf("Incorrect! The number is less than %d", guess)
	}

	fmt.Printf("\nGame over... You've run out of attempts.\nThe number was %d.\n", randomNumber)
}
