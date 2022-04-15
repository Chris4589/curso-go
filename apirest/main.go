package main

import (
	"dbmysql/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//nueva ruta
	mx := mux.NewRouter()

	//endpoints
	mx.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mx.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	// 	post
	mx.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")

	mx.HandleFunc("/api/user", handlers.UpdateUser).Methods("PUT")

	mx.HandleFunc("/api/user", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3000", mx))
}
