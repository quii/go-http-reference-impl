package http

import (
	"github.com/gorilla/mux"
	"github.com/quii/hellok8s/internal/domain/greet"
	"github.com/quii/hellok8s/internal/http/internal"
)

func newRouter() *mux.Router {
	greetingHandler := internal.NewGreetHandler(internal.GreeterFunc(greet.HelloGreeter))

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", greetingHandler.Greet)
	
	return router
}
