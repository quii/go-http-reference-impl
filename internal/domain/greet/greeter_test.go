package greet

import (
	acceptance_criteria "github.com/quii/hellok8s/acceptance-criteria"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	acceptance_criteria.GreetingCriteria(t, acceptance_criteria.GreetingSystemFunc(HelloGreeter))
}
