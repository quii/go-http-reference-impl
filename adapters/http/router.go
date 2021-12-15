package http

import (
	"github.com/gorilla/mux"

	"github.com/quii/go-http-reference-impl/adapters/http/handlers"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/handlers/healthcheck", handlers.HealthCheck)

	return router
}
