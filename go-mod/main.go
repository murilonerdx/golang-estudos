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
		Nome:             "Murilo P.S",
		Endereco:         endereco,
		DataDeNascimento: time.Date(2001, 04, 21, 0, 0, 0, 0, time.Local),
	}

	//fmt.Print(endereco)
	//endereco.Numero = 16
	//fmt.Print(endereco)
	//fmt.Print(endereco)
	fmt.Print(pessoa)
	//idade := model.CaculaIdade(pessoa)
	idade := pessoa.IdadeAtual()

	fmt.Print("Pessoa tem ", idade, " anos")

	//A importancia de usar ponteiros para alterar valores na memoria se dentro do struct deixar sem o * acaba não atribuindo
	model.CaculaIdade(pessoa)
	fmt.Print(pessoa)
}
