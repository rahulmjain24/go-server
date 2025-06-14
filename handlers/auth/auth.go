package auth

import (
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Registered"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User logged in"))
}
