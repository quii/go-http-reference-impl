package http

import (
	"net/http"

	"github.com/quii/go-http-reference-impl/application/ports"
)

func NewWebServer(
	config ServerConfig,
	greeter ports.GreeterService,
) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      newRouter(greeter),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
