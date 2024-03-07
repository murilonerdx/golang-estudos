package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	palavra := "Palavra teste"
	num := 0
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		num = i
		if i == 5 {
			break
		}
	}

	concatenated := strings.Join([]string{strconv.Itoa(num), palavra}, " ")

	fmt.Println(concatenated)
}
