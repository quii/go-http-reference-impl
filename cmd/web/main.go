package main

import (
	"context"
	"github.com/quii/go-http-reference-impl/internal/adapters/http"
	"log"
	"os"
	"os/signal"
)

const (
	exitCodeInterrupt = 2
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()

	app := newApp(ctx)

	server := http.NewWebServer(
		app.ServerConfig,
		app.Greeter,
		app.RecipeService,
	)

	log.Fatal(server.ListenAndServe())
}

