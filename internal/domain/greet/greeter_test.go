package greet

import (
	acceptance_tests "github.com/quii/hellok8s/acceptance-tests"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	acceptance_tests.GreetingAcceptanceTest(t, acceptance_tests.GreetingSystemFunc(HelloGreeter))
}
