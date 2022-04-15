package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	msg := make(map[string]string, 0)
	msg["msg"] = "hola mundo"

	http.HandleFunc("/params/", func(res http.ResponseWriter, req *http.Request) {
		// values := req.URL.Query().Get("id") // id=1 QUERYPARAMS
		values := req.URL.Query() //mapa de QUERY params (id=1 id2=2) todos juntos se usa .Get("value") para tener el deseado
		response, _ := json.Marshal(values)
		fmt.Println(values.Get("id"))
		fmt.Print("asdadssadasasdasdsddsd")
		//.Has tiene
		if values.Has("id") {
			fmt.Print("si tiene id\n")
		}

		res.Header().Add("Content-Type", "application/json") // json
		fmt.Fprintln(res, string(response))
	})

	//router
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		response, _ := json.Marshal(msg)
		res.Header().Add("Content-Type", "application/json") // json
		fmt.Fprintln(res, string(response))
	})

	http.HandleFunc("/notfound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	msg2 := map[string]string{
		"error": "server internal",
	}
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		response, _ := json.Marshal(msg2)
		r.Header.Add("Content-Type", "application/json")
		http.Error(w, string(response), http.StatusInternalServerError)
	})

	//crear sv para arrancar sv web
	fmt.Println("el servidor corre en el púerto http:///localhost:3000")

	//CREACION MUX
	mux := http.NewServeMux()
	//se usa ahora mux y ya no http y se envia en listenserver en lugar de nil
	//deja de funcionar los que no esten con mux.handle
	mux.HandleFunc("/mux", func(res http.ResponseWriter, req *http.Request) {
		response, _ := json.Marshal(msg)
		res.Header().Add("Content-Type", "application/json") // json
		fmt.Fprintln(res, string(response))
	})
	server := &http.Server{
		Addr:    "localhost:3000", //direccion
		Handler: mux,              //handler
	}
	// NIL = TRABAJA CON MUX AUTOMATICOS
	//log.Fatal(http.ListenAndServe("localhost:3000", mux)) //MUX
	log.Fatal(server.ListenAndServe())
}
