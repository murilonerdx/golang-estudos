package main

import "fmt"

func main() {
	ImprimiMensage("Hello funcao")
}

func ImprimiMensage(message string) {
	fmt.Print(message)

	message += " ,teste 123"
	fmt.Print(message)
}
