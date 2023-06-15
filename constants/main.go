package main

import "fmt"

func main() {
	// for int
	const number int = 5
	fmt.Printf("Data: %d, and it's type %T \n", number, number)

	// for string
	const hello string = "Hello World"
	fmt.Printf("Data: %s, it's type %T \n", hello, hello)

	// for float
	const fnun float64 = 2.5
	fmt.Printf("Data: %f, it's type %T\n", fnun, fnun)

	// for bool
	const boolean bool = true
	fmt.Printf("Data: %t, it's type %T\n", boolean, boolean)

}
