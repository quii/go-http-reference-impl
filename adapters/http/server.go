package http

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/quii/go-http-reference-impl/application/ports"
)

func NewWebServer(
	config ServerConfig,
	greeter ports.GreeterService,
) *http.Server {
	router := newRouter(greeter)
	return &http.Server{
		Addr: config.TCPAddress(),
		Handler: otelhttp.NewHandler(
			router,
			"greet-http-server",
		),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
