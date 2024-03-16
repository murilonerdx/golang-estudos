package main

import "fmt"

func main() {
	fmt.Print(OperacaoNomeada(1, 2))
}

func OperacaoNomeada(numero1 int, numero2 int) (soma int, multiplicao int, divisao int, subtracao int) {
	soma = numero1 + numero2
	multiplicao = numero1 * numero2
	divisao = numero1 / numero2
	subtracao = numero1 - numero2
	return
}
