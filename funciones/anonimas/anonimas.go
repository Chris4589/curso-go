package main

import "fmt"

func main() {

	func(rest ...int) {
		fmt.Println("hola desde la func")
	}()

	myFunc := func() {
		fmt.Println("holawe")
	}

	myFunc()
}
