package handlers

import (
	"fmt"
	"net/http"
	"rest/utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(funcion customeHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuthenticated(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}
		funcion(w, r)
	})
}

func MiddlewareTwo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Este es el segundo wrap")
		handler.ServeHTTP(w, r)
	})
}
