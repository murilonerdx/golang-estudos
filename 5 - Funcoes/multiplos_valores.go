package main

import "fmt"

func main() {
	fmt.Print(Operacao(1, 2))
}

func Operacao(numero1 int, numero2 int) (int, int, int, int, int) {
	resultado := numero1 + numero2
	return resultado, resultado * 2, resultado * 3, resultado * 4, resultado * 5
}
