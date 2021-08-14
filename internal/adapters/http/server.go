package http

import (
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http"
)

func NewWebServer(
	config ServerConfig,
	greeter ports.GreeterService,
	recipeService ports.RecipeService,
) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      newRouter(greeter, recipeService),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
