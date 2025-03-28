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
	sum := 0
	for i := 1; i <= num; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
