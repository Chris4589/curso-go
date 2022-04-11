package main

/* paquetes de terceros */
// go get <url>
import (
	creacionmodulo "github.com/Chris4589/module_test"
	"github.com/donvito/hellomod"
)

func main() {

	hellomod.SayHello()
	creacionmodulo.Hola()
}
