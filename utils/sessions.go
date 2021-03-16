package utils

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

const (
	cookieName   = "go_session"
	cookiExpires = 24 * 2 * time.Hour // dos dias
)

func SetSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   uuid.NewV4().String(),
		Path:    "/",
		Expires: time.Now().Add(cookiExpires),
	}
	http.SetCookie(w, cookie)
}

func DeleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		return cookie.Value // uuid
	}
	return ""
}

func IsAuthenticated(r *http.Request) bool {
	return getValCookie(r) != ""
}
