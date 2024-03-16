package main

import "fmt"

// Executa primeiro por ser init inicialização inicia antes do main()
func init() {
	fmt.Print("teste init")
}

// var numeroMagico = 5
func main() {
	fmt.Print(OperacaoNomeadaInit(1, 2))
}

// Cuidado com os escopos o Golang permiti usar variaveis dentro de funções que sobrescrever o valor no arquivo por exemplo no comentario
func OperacaoNomeadaInit(numero1 int, numero2 int) (soma int, multiplicao int, divisao int, subtracao int) {
	//numeroMagico := 3
	soma = numero1 + numero2
	multiplicao = numero1 * numero2
	divisao = numero1 / numero2
	subtracao = numero1 - numero2
	return
}
