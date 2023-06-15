package main

import "fmt"

func main() {
	var a int = 20
	var b int = 10

	fmt.Println(fmt.Sprintf("a = %d & b = %d", a, b))

	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)

}
