package models

type Persona2 struct {
	nombre string
	edad   int
}

func (p *Persona2) Constructor(nombre string, edad int) {
	p.edad = edad
	p.nombre = nombre
}

//usar getter y seter

func (p *Persona2) GetNombre() string {
	return p.nombre
}

func (p *Persona2) GetEdad() int {
	return p.edad
}

func (p *Persona2) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona2) SetEdad(edad int) {
	p.edad = edad
}
