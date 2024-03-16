package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func main() {
	endereco := endereco{
		rua:    "Rua x",
		numero: 15,
		cidade: "São Paulo",
	}

	fmt.Print(endereco)
	endereco.numero = 16
	fmt.Print(endereco)
	fmt.Print(endereco)

}
