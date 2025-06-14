package main

import (
	"net/http"

	"github.com/rahulmjain24/crud-server/api"
	"github.com/rahulmjain24/crud-server/config"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT", "8080")

	r := api.RegisterRoutes()

	http.ListenAndServe(":"+port, r)
}
