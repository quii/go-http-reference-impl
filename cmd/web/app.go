package main

import (
	"github.com/quii/go-http-reference-impl/internal/adapters/http"
	in_mem "github.com/quii/go-http-reference-impl/internal/adapters/in-mem"
	"github.com/quii/go-http-reference-impl/internal/domain/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
)

// App holds and creates dependencies for your application
type App struct {
	ServerConfig  http.ServerConfig
	Greeter       ports.GreeterService
	RecipeService ports.RecipeService
}

func newApp() *App {
	config := newDefaultConfig()
	return &App{
		ServerConfig:  config,
		Greeter:       ports.GreeterServiceFunc(greet.HelloGreeter),
		RecipeService: in_mem.NewRecipeStore(),
	}
}
