package hello_go_k8s

import (
	"github.com/matryer/is"
	"testing"
)

type GreetingSystem interface {
	Greet(name string) (greeting string, err error)
}

func GreetingAcceptanceTest(t *testing.T, system GreetingSystem) {
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := system.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

