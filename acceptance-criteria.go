package hello_go_k8s

import (
	"github.com/matryer/is"
	"testing"
)

func GreetingAcceptanceTest(t *testing.T, client *APIClient) {
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := client.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

