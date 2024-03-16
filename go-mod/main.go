package main

import (
	"fmt"
	"teste-go-mod/model"
)

var endereco model.Endereco

func main() {
	endereco = model.Endereco{
		Rua:    "Rua x",
		Numero: 15,
		Cidade: "São Paulo",
	}

	fmt.Print(endereco)
	endereco.Numero = 16
	fmt.Print(endereco)
	fmt.Print(endereco)

}
