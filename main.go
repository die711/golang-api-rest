package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/users/").Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}").Methods("GET")
	mux.HandleFunc("/api/v1/users/").Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}").Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}").Methods("DELETE")

	log.Println("El servidor ala esucha en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
