//go:build unit
// +build unit

package http_test

import (
	"net/http/httptest"
	"testing"

	"github.com/quii/go-http-reference-impl/internal/adapters/http"

	"github.com/quii/go-http-reference-impl/black-box-tests/acceptance"
	in_mem "github.com/quii/go-http-reference-impl/internal/adapters/in-mem"
	"github.com/quii/go-http-reference-impl/internal/domain/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"github.com/quii/go-http-reference-impl/specifications"
)

func TestNewWebServer(t *testing.T) {
	webServer := http.NewWebServer(
		http.ServerConfig{},
		ports.GreeterServiceFunc(greet.HelloGreeter),
		in_mem.NewRecipeStore(),
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := acceptance.NewAPIClient(svr.URL, t)

	specifications.Greeting(t, client)
}
