package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rahulmjain24/go-server/services"
)

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := services.RegisterUser(req.Name, req.Email)
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User logged in"))
}
