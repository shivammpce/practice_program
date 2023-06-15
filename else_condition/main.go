package main

import (
	"fmt"
)

func main() {
	var a int = 123

	if a > 100 {
		//	fmt.Println(fmt.Sprintf("number %d is greater than 100", a))
		info := fmt.Sprintf("number %d is greater than 100", a)
		fmt.Println(info)
	} else {
		info := fmt.Sprintf("number %d is less than 100", a)
		fmt.Println(info)
	}
}
