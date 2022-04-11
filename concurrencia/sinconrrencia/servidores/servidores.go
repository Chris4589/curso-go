package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
sin concurrencia el script revisa sv x sv 1x1 y tarda N tiempo hasta que al cola se ejecute totalmente
*/
func revisarServidor(servidor string) {
	//packete http
	response, error := http.Get(servidor)

	if error != nil {
		fmt.Println("servidor no disponible")
	} else {
		fmt.Println(response.StatusCode)
	}
}

func main() {
	inicio := time.Now()
	servidores := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/12",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/7",
		"https://jsonplaceholder.typicode.com/posts/9",
		"https://jsonplaceholder.typicode.com/posts",
	}

	for _, servidor := range servidores {
		//revisa 1x1
		//dura 2s
		//sin concurrencia
		revisarServidor(servidor)
	}
	end := time.Since(inicio)
	fmt.Println(end)
}
