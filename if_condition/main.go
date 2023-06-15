package main

import "fmt"

func main() {
	var a int = 10

	if a < 100 {
		//	fmt.Println(true)
		info := fmt.Sprintf("number %d is less than 100", a)
		//	fmt.Println(fmt.Sprintf("number %d is less than 100", a))
		fmt.Println(info)
	}
}
