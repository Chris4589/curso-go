package main

import (
	"fmt"
	"paquete/models"
)

//strutc son las clases de go
//serian el molde

//CREAR CLASE
//structura(clase)
type Person struct {
	//atributos
	edad   int
	nombre string
}

//metodos
func (p *Person) setNombre(name string) {
	p.nombre = name
}

//metodos
func (p *Person) getName() string {
	return p.nombre
}

//HERENCIA
type Empleado struct {

	//herencia
	//el extends de go
	Person
	sueldo float32
}

func (e *Empleado) getSueldo() float32 {
	return e.sueldo
}

func main() {
	//crear objetos desde la estructura
	pt := Person{1, "christopher"}

	fmt.Println(pt)
	pt.setNombre("david")
	fmt.Println(pt)

	//crear objeto 2
	//para no respetar el orden del constructor
	p2 := Person{
		nombre: "ch",
		edad:   24, //va , al final
	}
	fmt.Println(p2.getName())

	//objeto de tipo empleado que hereda cosas de persona
	emp := Empleado{
		Person: Person{
			edad:   50,
			nombre: "test emp",
		},
		sueldo: 3200,
	}

	fmt.Println(emp)
	fmt.Println(emp.getSueldo())

	pe1 := models.Persona2{}
	pe1.Constructor("asdsd", 26)
	fmt.Println(pe1.GetEdad())
	fmt.Println(pe1.GetNombre())
	fmt.Println(pe1)
	pe1.SetEdad(21)
	pe1.SetNombre("chris")
	fmt.Println(pe1)
}
