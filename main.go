package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/handlers"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Println("El servidor ala esucha en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
