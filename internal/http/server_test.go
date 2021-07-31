// +build unit

package http

import (
	"github.com/matryer/is"
	hello_go_k8s "github.com/quii/hellok8s"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := hello_go_k8s.NewAPIClient(server.URL)

	t.Run("healthcheck", func(t *testing.T) {
		is := is.New(t)
		is.NoErr(client.CheckIfHealthy())
	})

	t.Run("greeting", func(t *testing.T) {
		is := is.New(t)

		greeting, err := client.Greet()
		is.NoErr(err)
		is.Equal(greeting, "Hello, world!")
	})
}
