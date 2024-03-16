package main

import "fmt"

// Os ponteiros tem um motivo, o uso de memoria como acontece no java prejudica em grande escala o Garbage Collector, pois acaba utilizando muita memoria do sistema
func main() {
	x := 8
	// y := x não referencia mas aloca ao valor da memoria de x[x=0x105...] , y[0x105...]
	y := &x //& referencia ao valor da memoria de x,y [x=0x105...] há partir desse momento não conseguimos definir valores pois ele vira imutavel
	fmt.Println(y, &x)
	fmt.Println(&y, &x)
	*y = 10 // * serve justamente para atribuir um valor e transformar ele novamente em mutavel
	fmt.Println(y, &y)

	//Ao mudar o ponteiro *y = 10. x,y[0x105] você muda ambos x,y[0x016]
	fmt.Println(*y, x)

}
