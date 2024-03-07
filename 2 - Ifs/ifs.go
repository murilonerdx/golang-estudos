package main

import "fmt"

func main() {

	salario := 990.00
	var salarioMaisOBonus float64
	salarioMaisOBonus = salario
	var isCar bool = true

	for i := 10; i > 10; i-- {
		fmt.Println(i)
	}

	if salario < 1000 && isCar {
		fmt.Println("É um carro")
		salarioMaisOBonus = salarioMaisOBonus - 100
	} else {
		fmt.Println("Não é um carro")
		salarioMaisOBonus = salarioMaisOBonus + 100
	}

	fmt.Println("Salário: ", salarioMaisOBonus)

}
