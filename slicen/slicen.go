package main

import "fmt"

func main() {
	//slicen
	//son mutables
	numeros := []int{12, 2, 4}
	fmt.Println(numeros, len(numeros))

	//agregar elementos

	numeros = append(numeros, 1)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//sub slicen
	subSlicen := numeros[:2]

	numeros[0] = 1 // si el padre cambia al ser un slice tambien afecta al hijo

	fmt.Println(subSlicen, len(subSlicen))
	fmt.Println(numeros, len(numeros))

	//capacidad
	//púnteros
	//lñongitud
	meses := []string{"enero", "febrero"}

	fmt.Printf("len %v cap %v %p", len(meses), cap(meses), meses)

	meses = append(meses, "marzo")

	fmt.Printf("len %v cap %v %p", len(meses), cap(meses), meses)

}
