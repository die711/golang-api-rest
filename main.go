package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/config"
	"rest/handlers"
	"rest/handlers/api/v1"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/users/new", handlers.NewUser).Methods("GET", "POST")

	mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", v1.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")

	log.Println("El servidor ala esucha en el puerto ", config.ServerPort())

	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}
