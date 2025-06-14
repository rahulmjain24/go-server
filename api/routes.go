package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rahulmjain24/go-server/handlers"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.HelloHandler)

	r.Mount("/auth", AuthRoutes())

	return r
}
