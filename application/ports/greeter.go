package ports

import "context"

//go:generate moq -out greeterservice_moq.go . GreeterService
type GreeterService interface {
	Greet(ctx context.Context, name string) (greeting string, err error)
}

type GreeterServiceFunc func(context.Context, string) (string, error)

func (g GreeterServiceFunc) Greet(ctx context.Context, name string) (greeting string, err error) {
	return g(ctx, name)
}
