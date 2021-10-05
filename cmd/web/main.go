package main

import (
	"github.com/quii/go-http-reference-impl/internal/adapters/http"
	"log"
)

func main() {
	app := newApp()

	server := http.NewWebServer(
		app.ServerConfig,
		app.Greeter,
		app.RecipeService,
	)

	log.Fatal(server.ListenAndServe())
}

