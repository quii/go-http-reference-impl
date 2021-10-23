package specifications

import (
	"testing"

	"github.com/matryer/is"
)

type GreetingSystemDriver interface {
	Greet(name string) (greeting string, err error)
}

func Greeting(t *testing.T, greetingSystem GreetingSystemDriver) {
	t.Helper()
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := greetingSystem.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

type GreetingSystemFunc func(name string) (greeting string, err error)

func (g GreetingSystemFunc) Greet(name string) (greeting string, err error) {
	return g(name)
}
