package main

import (
	"fmt"
	"os"
)

//funciones

//defer
/*
	definicion
	se ejecuta al final de la funcion sin importar donde este declarado

	panic: error garrafal

*/

//panic
/*
	definicion

	cierra la app de manera brusca
	---
	goroutine 1 [running]:
	main.main()
					C:/Users/Casa/go/src/curso/readfile/readfile.go:35 +0x1ce
	exit status 2
	---
*/

//recover
/*
	definicion

	controla todos los panicos

	se usa por si se escapa un panic (excepcion) en la app ya que no se puede controlar o no
	se sabe de donde viene el error, para que no se cierre de manera brusca se usa "recover"
*/

func main() {

	//recover
	//se agarran los errores de panic para no cerrar el error de forma brusca
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("el programa no finalizo de forma correcta")
		}
	}()

	//variable tipo File
	file, error := os.Open("C:2\\Users\\Casa\\go\\src\\curso\\readfile\\hola.txt")

	if error != nil {
		fmt.Println(error)
		panic("error garrafal") //se cierra el programa de forma brusca

	} else {
		defer func() {
			fmt.Println("cerro")
			file.Close()
		}() //auto invocada ()

		//viene en bytes
		contenido := make([]byte, 254) //slice
		// contenido := []byte{}
		longitud, _ := file.Read(contenido)
		contenidoFile := string(contenido[:longitud]) //0 to longitud
		fmt.Println(contenidoFile)
	}

	fmt.Println("antes de cerrar")
}
