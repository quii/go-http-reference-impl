package main

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/quii/go-http-reference-impl/adapters/http"
)

// App holds and creates dependencies for your application.
type App struct {
	ServerConfig http.ServerConfig
}

func newApp(applicationContext context.Context) (*App, error) {
	config, err := newDefaultConfig()
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(config.TraceProvider)

	go cleanupBeforeQuit(applicationContext, config.TraceProvider)

	return &App{
		ServerConfig: config,
	}, nil
}

// here is a chance to tidy things up before the app quits.
func cleanupBeforeQuit(ctx context.Context, tp *trace.TracerProvider) {
	<-ctx.Done()
	log.Println("☠️  Program has been told to quit, I should tidy things up ☠️")

	if err := tp.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("✅ Shutdown tracer provider")
	os.Exit(0) // not sure if this is "good"
}
