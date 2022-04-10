package main

import (
	"fmt"
)

var mensaje string

//variables globales
func change() {
	mensaje = "change"
	fmt.Println(mensaje)
}

func main() {
	change()
	mensaje = "main"
	fmt.Println(mensaje)
}
