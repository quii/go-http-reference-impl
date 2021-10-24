package greet_test

import (
	"testing"

	"github.com/quii/go-http-reference-impl/application/greet"

	"github.com/quii/go-http-reference-impl/specifications"
)

func TestHelloGreeter(t *testing.T) {
	helloGreeter := specifications.GreetingSystemFunc(greet.HelloGreeter)
	specifications.Greeting(t, helloGreeter)
}
