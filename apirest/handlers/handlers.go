package handlers

import (
	"dbmysql/db"
	"dbmysql/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db.Connect()
	users := models.ListUsers()
	db.Close()

	out, _ := json.Marshal(users)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(rw, string(out))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	db.Connect()

	//obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	users := models.GetUser(userId)
	db.Close()

	out, _ := json.Marshal(users)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(rw, string(out))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	//obtener registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body) //obtener el body y decodificar de json a objetoGO

	if decoder.Decode(&user) != nil { // lo lee lo decodificado y lo almacena en ( &parametro )
		fmt.Fprintln(rw, http.StatusInternalServerError)
	} else {
		db.Connect()
		user.Save() //lo guarda
		db.Close()

		out, _ := json.Marshal(user)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, string(out))
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//obtener registro
	// user := models.User{}
	db.Connect()
	query := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(query)

	user := models.GetUser(id)

	if user.Id == 0 {
		fmt.Fprintln(rw, http.StatusNoContent)
	} else {
		decoder := json.NewDecoder(r.Body) //obtener el body y decodificar de json a objetoGO

		if decoder.Decode(&user) != nil { // lo lee lo decodificado y lo almacena en ( &parametro )
			fmt.Fprintln(rw, http.StatusInternalServerError)
		} else {

			user.Save() //lo actualiza

			out, _ := json.Marshal(user)
			rw.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(rw, string(out))
		}
	}

	db.Close()
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	db.Connect()
	query := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(query)

	user := models.GetUser(id)

	if user.Id == 0 {
		fmt.Fprintln(rw, http.StatusNoContent)
	} else {
		decoder := json.NewDecoder(r.Body) //obtener el body y decodificar de json a objetoGO

		if decoder.Decode(&user) != nil { // lo lee lo decodificado y lo almacena en ( &parametro )
			fmt.Fprintln(rw, http.StatusInternalServerError)
		} else {

			user.Delete() //lo actualiza

			out, _ := json.Marshal(user)
			rw.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(rw, string(out))
		}
	}

	db.Close()
}
