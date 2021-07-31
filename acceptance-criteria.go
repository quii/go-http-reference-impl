package hello_go_k8s

import (
	"github.com/matryer/is"
	"testing"
)

func APIAcceptanceTest(t *testing.T, client *APIClient) {
	t.Run("healthcheck", func(t *testing.T) {
		is := is.New(t)
		is.NoErr(client.CheckIfHealthy())
	})

	t.Run("greeting", func(t *testing.T) {
		is := is.New(t)

		greeting, err := client.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

