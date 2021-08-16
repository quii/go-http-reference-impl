package greet

import (
	"github.com/quii/go-http-reference-impl/internal/acceptance_criteria"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	acceptance_criteria.GreetingCriteria(t, acceptance_criteria.GreetingSystemFunc(HelloGreeter))
}
