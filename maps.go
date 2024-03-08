package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["cpf"] = "999999999"
	m["cidade"] = "São Paulo"
	m["cep"] = "019273000"
	m["uf"] = "SP"

	//fmt.Print(m)

	var chave = "cidade"
	//valor, existe := m["cidade"]

	//if existe {
	//	fmt.Println(valor)
	//} else {
	//	fmt.Println("Chave não existe!")
	//}

	if valor, existe := m[chave]; existe {
		fmt.Println("chave:[", chave, "] valor:[", valor, "]")
	} else {
		fmt.Println("Chave não existe!")
	}

	for c, v := range m {
		fmt.Println("chave:[", c, "] valor:[", v, "]")
	}

}
