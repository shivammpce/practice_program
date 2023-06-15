package main

// 1
import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// 2
func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(1).Seconds().Do(func() {
		hello("John Doe")
	})

	// 5
	s.StartBlocking()
}

// 6
func main() {
	runCronJobs()
}

// The above code accomplishes the following:

// Imports the required packages
// Defines the hello function that prints a message and takes a string parameter called name
// Creates a new scheduler instance s in the runCronJobs function using the gocron package
// Defines a cron job that runs every one second (Every(1).Seconds()) and calls the hello function
// Starts the scheduler in blocking mode (StartBlocking()), which blocks the current execution path
// Calls the runCronJobs function inside the main function
