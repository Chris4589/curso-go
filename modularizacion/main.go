package main

import "paquetes/mensajes"

//modularizaicon
//un packete es una carpeta

//go mod init <name>
//go mod init paquetes

func main() {
	mensajes.Hola()
	mensajes.ExecutePrivate()
}
