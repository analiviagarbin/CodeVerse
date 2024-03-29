package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns the router with the configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
