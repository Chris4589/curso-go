package main

import "fmt"

//sirve para molimorfismo

//manejan los metodos que son lo mismo de las estructuras
// y lo mestran de manera dif
//maneja los metodos identicos de las estructuras
// y muestra algo distinto por cada metodo

type Animal interface {
	mover() string
}

type Perro struct {
}

type Pez struct {
}

type Pajaro struct {
}

func (*Perro) mover() string {
	return "camina"
}

func (*Pez) mover() string {
	return "nada"
}

func (*Pajaro) mover() string {
	return "vuela"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {
	perro := Perro{}

	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}

	moverAnimal(&pajaro)
}
