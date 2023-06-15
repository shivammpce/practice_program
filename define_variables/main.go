package main

import "fmt"

func main() {
	// for string
	var string_data string
	string_data = "Golang"
	fmt.Printf("Data: %s and it's type %T \n", string_data, string_data)

	// for int
	var int_data int
	int_data = 5
	fmt.Printf("Data: %d and it's type %T \n", int_data, int_data)

	// for float
	var float_data float64
	float_data = 0.5
	fmt.Printf("Data: %f and it's type %T\n", float_data, float_data)

	// for bool

	var bool_data bool
	bool_data = true
	fmt.Printf("Data: %t it's type %T \n", bool_data, bool_data)
}
