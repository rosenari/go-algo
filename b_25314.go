package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(input))

	bytes := num / 4
	for i := 1; i <= bytes; i++ {
		fmt.Print("long ")
	}
	fmt.Println("int")
}
