package main

import "fmt"

func CombOfNum(arr []int) []int {

	//arr := []int{1, 2, 3}

	for i := range arr {
		for j := range arr {
			for k := range arr {
				fmt.Println(arr[i], arr[j], arr[k])
			}
		}
	}
	return arr
}

func main() {
	CombOfNum([]int{1, 2, 3})
}
