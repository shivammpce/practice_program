package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 17
	if isPrime(n) {
		fmt.Printf("%d is a prime number\n", n)
	} else {
		fmt.Printf("%d is not a prime number\n", n)
	}
}
