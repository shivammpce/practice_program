package main

import "fmt"

func main() {
	var a string = "Golang"
	var p *string = &a

	fmt.Println("address of a: ", &a)
	fmt.Println("pointer p: ", p)
	fmt.Println("pointer p pointed to date: ", *p)

}
