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

	pessoa := model.Pessoa{
		"Murilo P.S",
		endereco,
	}

	//fmt.Print(endereco)
	//endereco.Numero = 16
	//fmt.Print(endereco)
	//fmt.Print(endereco)
	fmt.Print(pessoa)

}
