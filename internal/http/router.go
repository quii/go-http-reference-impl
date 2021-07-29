package http

import (
	"github.com/gorilla/mux"
	"github.com/quii/hellok8s/internal/http/handlers"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", handlers.HealthCheck)
	return router
}
