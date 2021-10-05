package main

import (
	"context"
	"github.com/quii/go-http-reference-impl/internal/adapters/http"
	in_mem "github.com/quii/go-http-reference-impl/internal/adapters/in-mem"
	"github.com/quii/go-http-reference-impl/internal/application/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"log"
)

// App holds and creates dependencies for your application
type App struct {
	ServerConfig  http.ServerConfig
	Greeter       ports.GreeterService
	RecipeService ports.RecipeService
}

func newApp(applicationContext context.Context) *App {
	config := newDefaultConfig()

	go doSomethingOnInterrupt(applicationContext)

	return &App{
		ServerConfig:  config,
		Greeter:       ports.GreeterServiceFunc(greet.HelloGreeter),
		RecipeService: in_mem.NewRecipeStore(),
	}
}

// this is just an example of how the services within an app could listen to the
// cancellation signal and stop their work gracefully. So it's probably a decent
// idea to pass the application context to services if you want to do some
// cleanup before finishing
func doSomethingOnInterrupt(ctx context.Context) {
	<-ctx.Done()
	log.Println("☠️ Program has been told to quit, I should tidy things up ☠️")
}
