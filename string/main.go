package main

import (
	"fmt"
	"strconv"
)

func main() {

	// m := make(map[string]int)
	r := "100fgfg564hdfbhfbf365dbfdbf"
	temp := []string{}

	a := []rune(r)
	for j, val := range a {
		num := ""
		_, err := strconv.Atoi(string(val))
		if err == nil {

			if j > 0 {
				_, err = strconv.Atoi(string(r[j-1]))
				if err == nil || len(num) == 0 {
					num += string(val)
				}
			} else {
				num += string(val)
			}
		}
		temp = append(temp, num)

	}
	fmt.Println(temp)

	// for i , val := range a {
	// 	fmt.Println("val = ", string(val))
	// 	mval , ok := m[string(val)]
	// 	fmt.Println("mval = ", string(mval) , " ok = ", ok)
	// 	if ok {
	// 		if mval == i-1 {
	// fmt.Println(fres)
	// 			break
	// 		}
	// 	}else{
	// 		m[string(val)] = i
	// 	}
	// 	fmt.Println("map = ", m)

	// 	// b := fmt.Sprintf("val = %s",string(val))
	// 	// fmt.Println(b)
	// }

}
