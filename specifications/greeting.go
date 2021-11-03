package specifications

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

type GreetingSystemDriver interface {
	Greet(ctx context.Context, name string) (greeting string, err error)
}

func Greeting(t *testing.T, greetingSystem GreetingSystemDriver) {
	t.Helper()
	t.Run("greets people in a friendly manner", func(t *testing.T) {
		is := is.New(t)

		greeting, err := greetingSystem.Greet(context.Background(), "Pepper")
		is.NoErr(err)
		is.Equal(greeting, "Hello, Pepper!")
	})
}

type GreetingSystemFunc func(ctx context.Context, name string) (greeting string, err error)

func (g GreetingSystemFunc) Greet(ctx context.Context, name string) (greeting string, err error) {
	return g(ctx, name)
}
