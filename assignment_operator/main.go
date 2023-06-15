package main

import "fmt"

func main() {
	var a int = 20
	var b int = 10
	var c int

	fmt.Println(fmt.Sprintf("value of a: %d and value of b: %d", a, b))

	c = a + b
	fmt.Println("c = a+b : ", c)
	fmt.Println(fmt.Sprintf("value of a: %d and value of c: %d", a, c))

	c += a
	fmt.Println("c = c+a: ", c)
	fmt.Println(fmt.Sprintf("value of a: %d and value of c: %d", a, c))

	c += b
	fmt.Println("c = c+b: ", c)
	fmt.Println(fmt.Sprintf("value of b: %d and value of c: %d", b, c))

	c -= a
	fmt.Println("c = c-a: ", c)
	fmt.Println(fmt.Sprintf("value of a: %d and value of c: %d", a, c))

}
