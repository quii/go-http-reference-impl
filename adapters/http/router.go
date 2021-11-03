package http

import (
	"github.com/gorilla/mux"

	"github.com/quii/go-http-reference-impl/application/ports"

	"github.com/quii/go-http-reference-impl/adapters/http/internal"
	"github.com/quii/go-http-reference-impl/adapters/http/internal/greethandler"
)

func newRouter(greeter ports.GreeterService) *mux.Router {
	greetingHandler := greethandler.New(greeter)

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", greetingHandler.Greet)

	return router
}
