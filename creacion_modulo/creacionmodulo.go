package creacionmodulo

import "fmt"

//debe ser en mayuscula
//porque con minuscula se convierte en privada

const mensaje = "hola desde const"

//public
func Hola() {
	fmt.Println("hola desde packete msg")
}

//privada
func hola() {
	fmt.Println(mensaje)
}

func ExecutePrivate() {
	hola()
}
