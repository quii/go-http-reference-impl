//go:build unit
// +build unit

package http_test

import (
	"net/http/httptest"
	"testing"

	"github.com/quii/go-http-reference-impl/application/ports"

	http2 "github.com/quii/go-http-reference-impl/adapters/http"
	in_mem "github.com/quii/go-http-reference-impl/adapters/in-mem"
	"github.com/quii/go-http-reference-impl/application/greet"
	"github.com/quii/go-http-reference-impl/specifications"
)

func TestNewWebServer(t *testing.T) {
	webServer := http2.NewWebServer(
		http2.ServerConfig{},
		ports.GreeterServiceFunc(greet.HelloGreeter),
		in_mem.NewRecipeStore(),
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := http2.NewAPIClient(svr.URL, t)

	specifications.Greeting(t, client)
}
