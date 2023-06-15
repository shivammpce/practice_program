package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Print("Enter value of a : ")
	fmt.Scanln(&a)

	fmt.Print("Enter value of b : ")
	fmt.Scanln(&b)

	fmt.Print("Enter value of c : ")
	fmt.Scanln(&c)

	arr := []int{a, b, c}

	for i := range arr {
		for j := range arr {
			for k := range arr {
				fmt.Println(arr[i], arr[j], arr[k])
			}
		}
	}
}
