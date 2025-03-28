package main

import (
	"fmt"
)

func main() {
	check := make([]bool, 30)

	for i := 0; i < 28; i++ {
		var num int
		_, err := fmt.Scan(&num) // bufio.NewReader를 사용하면 에러남.
		if err != nil {
			fmt.Println(err)
		}
		check[num-1] = true
	}

	for i := 0; i < 30; i++ {
		if !check[i] {
			fmt.Println(i + 1)
		}
	}
}
