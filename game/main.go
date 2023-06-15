package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 10
	secretNumber := rand.Intn(10) + 1

	// Initialize variables
	var guess int
	var attempts int

	// Loop until the user guesses the correct number
	for {
		// Ask the user to guess a number
		fmt.Print("Guess a number between 1 and 10: ")
		fmt.Scan(&guess)

		// Increment the number of attempts
		attempts++

		// Check if the guess is correct
		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		} else {
			fmt.Println("Wrong guess. Try again.")
		}
	}
}
