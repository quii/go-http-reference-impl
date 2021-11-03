package main

import (
	"log"

	"github.com/quii/go-http-reference-impl/adapters/http"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	app := newApp(ctx)

	server := http.NewWebServer(
		app.ServerConfig,
		app.Greeter,
	)

	log.Println("All services started. App is ready! 🚀🚀🚀")

	if err := server.ListenAndServe(); err != nil {
		panic(err) // this is preferable to log.Fatal as we want our defer to run
	}
}
