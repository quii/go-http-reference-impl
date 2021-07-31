package http

import (
	"github.com/gorilla/mux"
	"github.com/quii/hellok8s/internal/http/internal"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", internal.Greet)
	return router
}
