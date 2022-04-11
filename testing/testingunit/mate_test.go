package testingunit

/*
	CADA FUNCION DEL TEST DETE LLEVAR "TEST" AL INICIO PARA QUE FUNCIONE
	go test -cover

	para saber el coverage


	//generar el coverage
	go test -coverprofile=coverage

	verlo en cmd el coverage por fn
	go tool cover -func=coverage



	verlo en html el coverage por fn
	go tool cover -html=coverage


	//para saber donde tarda tantop tiempo
	go test -cpuprofile=cou
	go tool pprof=cou
	top


	web genera en pdf o sgv html
	se necesita graphviz
*/

import (
	"fmt"
	"testing"
)

/* func TestSuma(t *testing.T) {
	suma := Suma(1, 2)

	if suma != 3 {
		t.Error("suma incorrecta se esperaba de resultado 3")
	}
}
*/
//mejor manera por tabla
func TestSumaMejorada(t *testing.T) {
	//se crea lista de estructuras

	//tabla := []struct { lo que lleva la estructura } { ...los items del arreglo de la estructura }
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("suma incorrecta se esperaba de resultado %d", item.c)
		}
	}
}

func TestNumeroMayorA(t *testing.T) {
	//se crea lista de estructuras

	//tabla := []struct { lo que lleva la estructura } { ...los items del arreglo de la estructura }
	tabla := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 4, 5},
	}
	for _, item := range tabla {
		total := GetMax(item.a, item.b)

		if total != item.c {
			t.Errorf("el mayor era %d", item.c)
		}
	}
}

func TestNumeroMayorB(t *testing.T) {
	//se crea lista de estructuras
	fmt.Print("we")
	//tabla := []struct { lo que lleva la estructura } { ...los items del arreglo de la estructura }
	tabla := []struct {
		a int
		b int
		c int
	}{
		{2, 4, 4},
		{4, 5, 5},
	}
	for _, item := range tabla {
		total := GetMax(item.a, item.b)

		if total != item.c {
			t.Errorf("el mayor era %d", item.c)
		}
	}
}
