package main

import (
	"log"

	"github.com/quii/go-http-reference-impl/adapters/http"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	app, err := newApp(ctx)
	if err != nil {
		panic(err)
	}

	server := http.NewWebServer(
		app.ServerConfig,
	)

	log.Println("All services started. App is ready! 🚀🚀🚀")

	if err := server.ListenAndServe(); err != nil {
		panic(err) // this is preferable to log.Fatal as we want our defer to run
	}
}
