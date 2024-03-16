package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		for b := 0; b <= 10; b++ {
			fmt.Println(fmt.Sprintf("%d X %d = %d", i, b, i*b))
		}
	}

	lista := make([]int, 1)
	lista = append(lista, 2)
	lista = append(lista, 3)

	for i := 0; i < len(lista); i++ {
		fmt.Println(lista[i])
	}

	var listaMaior = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	var listaPar = make([]int, 0)

	for a := 0; a < len(listaMaior); a++ {
		if listaMaior[a]%2 == 0 {
			listaPar = append(listaPar, listaMaior[a])
		}
	}

	fmt.Println("Pares ", listaPar)

}
