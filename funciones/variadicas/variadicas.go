package main

import "fmt"

// recibe N parametros
func sumar(numeros ...int) []int {
	fmt.Println(numeros)

	return numeros
}

//parametros normales van antes
//retornar 2 parametros
func retorno(nombre string, numeros ...int) (string, int) {
	return nombre, 1
}

func main() {
	sumar(1, 2, 3, 4, 6, 7, 8, 3)

	//leer los valores retornados
	nombre, _ := retorno("awa", 1)

	fmt.Println(nombre)
}
