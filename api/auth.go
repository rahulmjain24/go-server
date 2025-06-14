package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rahulmjain24/crud-server/handlers/auth"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/register", auth.RegisterHandler)
	r.Get("/login", auth.RegisterHandler)

	return r
}
