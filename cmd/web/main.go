package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/exporters/jaeger"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/quii/go-http-reference-impl/adapters/http"
)

func main() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

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

const (
	service     = "greeter"
	environment = "production"
	id          = 1
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}
