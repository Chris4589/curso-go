package main

import (
	"fmt"
	"strings"
)

//closures
func repeat(num int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, num)
	}
}

func main() {
	repeats := repeat(3)("we")

	fmt.Println(repeats)
}
