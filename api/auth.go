package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rahulmjain24/go-server/handlers"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/register", handlers.RegisterHandler)
	r.Get("/login", handlers.LoginHandler)

	return r
}
