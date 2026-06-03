package main

import (
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

	for {
		fmt.Println("\nPlease select the difficulty level: ")
		fmt.Println("1 - Easy (10 chances)")
		fmt.Println("2 - Medium (5 chances)")
		fmt.Println("3 - Hard (3 chances)")
		fmt.Println("4 - Quit")

		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

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
		case 4:
			fmt.Println("Quitting... Bye!")
			return
		default:
			fmt.Println("Please enter a valid number.")
			continue
		}

		fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game\n", difficultyLevel)

		var guess int
		for i := 0; i < numberOfAttempts; i++ {
			fmt.Print("\nEnter your guess: ")
			fmt.Scanln(&guess)

			if guess == randomNumber {
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", i+1)
				break
			}

			if i == numberOfAttempts-1 {
				fmt.Printf("\nGame over... You've run out of attempts.\nThe number was %d.\n", randomNumber)
				break
			}

			if guess < randomNumber {
				fmt.Printf("Incorrect! The number is greater than %d", guess)
				continue
			} else {
				fmt.Printf("Incorrect! The number is less than %d", guess)
				continue
			}

		}

		for {
			fmt.Print("\nWould you like to play again? (y/n) ")
			answer := ""
			fmt.Scanln(&answer)

			if answer == "y" || answer == "yes" {
				fmt.Println("Welcome to the Number Guessing Game! ")
				fmt.Println("I'm thinking of a number between 1 and 100")
				randomNumber = rand.Intn(101)
				break
			} else if answer == "n" || answer == "no" {
				fmt.Println("Quitting... Bye!")
				return
			} else {
				fmt.Println("Invalid input. Please enter y or n.")
			}
		}

	}
}
