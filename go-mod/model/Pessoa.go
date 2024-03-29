package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func CaculaIdade(p Pessoa) int {
	anoNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()

	return anoAtual - anoNascimento
}

func (p *Pessoa) IdadeAtual() int {
	anoNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()

	p.Idade = anoAtual - anoNascimento
	return anoAtual - anoNascimento
}
