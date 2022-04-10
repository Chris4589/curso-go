package main

import "fmt"

func modificar(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "wewewe"
}

func main() {
	cadena := "asdasd"
	modificar(&cadena)
	fmt.Printf("%p\n", &cadena)
}
