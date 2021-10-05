package greet

import (
	"github.com/quii/go-http-reference-impl/specifications"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	specifications.Greeting(t, specifications.GreetingSystemFunc(HelloGreeter))
}
