package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var somaAte5 = 0
	var somaAte10 = 0

	for i := 0; i < len(lista); i++ {
		if i < 5 {
			somaAte5 += lista[i]
		} else {
			somaAte10 += lista[i]
		}
	}

	fmt.Println("Soma até (5) ", somaAte5, " Soma até (10) ", somaAte10)

}
