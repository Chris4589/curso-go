package main

import "fmt"

func main() {
	//slicen
	//varios y tipo
	// tamaño
	//capacidad
	numeros := make([]int, 3, 3) // definir slicen vacio
	// numeros := make([]int, 0, 3)

	// fmt.Println(numeros, len(numeros), cap(numeros))

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
