package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene un usuario")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario")
}
