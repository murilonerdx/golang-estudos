package main

import (
	"fmt"
	"teste-go-mod/model"
	"time"
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
		time.Date(2001, 04, 21, 0, 0, 0, 0, time.Local),
	}

	//fmt.Print(endereco)
	//endereco.Numero = 16
	//fmt.Print(endereco)
	//fmt.Print(endereco)
	fmt.Print(pessoa)
	//idade := model.CaculaIdade(pessoa)
	idade := pessoa.IdadeAtual()

	fmt.Print("Pessoa tem ", idade, " anos")

}
