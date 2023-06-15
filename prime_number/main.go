package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 2; i <= 20; i++ {
		if isPrime(i) {
			fmt.Println(i, "is a prime number")
		}
	}
}
