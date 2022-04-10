package main

import "fmt"

func main() {
	a := 4
	b := 5

	r := 4 == 5

	// if
	if r {
		fmt.Printf("%d es igual a %d\n", a, b)
	} else {
		fmt.Printf("%d es no igual a %d \n", a, b)
	}

	// for
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}
