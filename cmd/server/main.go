package main

import (
	"net/http"

	"github.com/rahulmjain24/go-server/api"
	"github.com/rahulmjain24/go-server/config"
	"github.com/rahulmjain24/go-server/db"
)

func main() {
	config.LoadEnv()
	db.Init()

	port := config.GetEnv("PORT", "8080")

	r := api.RegisterRoutes()

	http.ListenAndServe(":"+port, r)
}
