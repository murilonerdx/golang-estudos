package main

import "fmt"

func main() {
	var lista = []int{1, 2, 5, 1, 2, 20, 20}
	var valueTotal = 0

	for i := 0; i < len(lista); i++ {
		valueTotal += lista[i]
	}

	fmt.Println("soma total ", valueTotal)

}
