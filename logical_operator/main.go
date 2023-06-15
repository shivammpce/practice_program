package main

import "fmt"

func main() {
	var a bool = false
	var b bool = true
	fmt.Println(fmt.Sprintf("a = %t & b = %t", a, b))

	fmt.Println("a && b: ", a && b)
	fmt.Println("a || b: ", a || b)
	fmt.Println("inverse of a: ", !a)
	fmt.Println("inverse of b: ", !b)
}
