package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	con concurrencia

	- GoRuntimes

	se ejecutaran todas al mismo tiempo

	SE EJECUTAN EN PARALELO


	- Channels
	se necesitan canales para saber los resultados y la ejecucion
*/
func revisarServidor(servidor string, channel chan string) {
	//packete http
	response, error := http.Get(servidor)

	if error != nil {
		// fmt.Println("servidor no disponible")
		// <- le envia lo que pasa al canal
		channel <- servidor + "servidor no disponible"
	} else {
		// fmt.Println(response.StatusCode)
		channel <- servidor + response.Status
	}
}

func main() {

	//cnales
	channels := make(chan string)

	inicio := time.Now()
	servidores := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/12",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://jsonplaceholder.typicode.com/posts",
	}

	for _, servidor := range servidores {
		// ACA SE EJECUTAN TODOS AL MISMO TIEMPO

		// go crea multiples hilos al mismo tiempo
		go revisarServidor(servidor, channels)
		fmt.Println(<-channels)
		//se necesitan canales para saber los resultados y la ejecucion
	}

	/* for i := 0; i < len(servidores); i++ {
		//aca se resuelve lo que devuelve  el canal
		fmt.Println(<-channels)
	} */

	end := time.Since(inicio)
	fmt.Println(end)
}
