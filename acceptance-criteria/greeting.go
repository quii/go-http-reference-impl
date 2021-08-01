package acceptance_criteria

import (
	"github.com/matryer/is"
	"testing"
)

type GreetingSystemAdapter interface {
	Greet(name string) (greeting string, err error)
}

func GreetingCriteria(t *testing.T, system GreetingSystemAdapter) {
	t.Helper()
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := system.Greet("Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

type GreetingSystemFunc func(name string) (greeting string, err error)

func (g GreetingSystemFunc) Greet(name string) (greeting string, err error) {
	return g(name)
}
