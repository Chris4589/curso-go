package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se puede dividir entre 0")
	}
	return dividendo / divisor, nil //es es cuando es un valor ausente
}

func main() {

	result, error := division(4, 0)

	if error != nil {
		fmt.Println("se envio 0")
	} else {
		fmt.Println(result)
	}
	//controlar una excepcion

}
