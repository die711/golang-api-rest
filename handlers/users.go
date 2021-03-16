package handlers

import (
	"net/http"
	"rest/utils"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "users/new", nil)
}
