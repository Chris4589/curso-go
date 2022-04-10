package main

import "fmt"

func main() {

	//a este si se le pueve añadiar desde pocisiones 0 1 2 3 etc, al slice no
	//[clave]valor
	dias := make(map[int]string) //mapa
	fmt.Println(dias)

	//agregar datos

	dias[1] = "lunes"
	dias[2] = "martes"
	dias[3] = "miercoles"
	dias[4] = "jueves"
	dias[5] = "viernes"
	dias[6] = "sabado"

	fmt.Println(dias)

	//modificar
	dias[5] = "VIENRS"
	dias[7] = "DOMINGO"

	fmt.Println(dias)

	//ELIMINAR

	delete(dias, 1)

	fmt.Println(dias, len(dias)) //adios lunes

	//nuevo mapa mas complejo

	estudiantes := make(map[string][]int)

	estudiantes["alex"] = []int{13, 15, 17}

	fmt.Println(estudiantes)

}
