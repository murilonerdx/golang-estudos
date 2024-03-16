package main

import "fmt"

func main() {

	var listaMaior = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	var listaPar = make([]int, 0)

	for a := 0; a < len(listaMaior); a++ {
		if listaMaior[a]%2 == 0 {
			listaPar = append(listaPar, listaMaior[a])
		}
	}

	fmt.Println("Pares ", listaPar)
	fmt.Println("Pares ", listaPar[:3])
	fmt.Println("Pares ", listaPar[3:])
	fmt.Println("Pares ", listaMaior[len(listaPar):])
	fmt.Println("Pares ", listaMaior[:len(listaPar)])

}
