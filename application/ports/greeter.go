package ports

//go:generate moq -out greeterservice_moq.go . GreeterService
type GreeterService interface {
	Greet(name string) (greeting string, err error)
}

type GreeterServiceFunc func(string) (string, error)

func (g GreeterServiceFunc) Greet(name string) (greeting string, err error) {
	return g(name)
}
