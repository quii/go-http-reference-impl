// +build unit

package http

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"github.com/quii/go-http-reference-impl/internal/specifications/adapters"
	in_mem "github.com/quii/go-http-reference-impl/internal/adapters/in-mem"
	"github.com/quii/go-http-reference-impl/internal/application/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {

	webServer := NewWebServer(
		ServerConfig{},
		ports.GreeterServiceFunc(greet.HelloGreeter),
		in_mem.NewRecipeStore(),
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := adapters.NewAPIClient(svr.URL, t)

	specifications.Greeting(t, client)
}
