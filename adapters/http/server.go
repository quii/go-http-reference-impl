package http

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewWebServer(
	config ServerConfig,
) *http.Server {
	router := newRouter()
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
