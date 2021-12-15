//go:build acceptance
// +build acceptance

package acceptance_test

import (
	"testing"

	"github.com/quii/go-http-reference-impl/adapters/http"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	client := http.NewAPIClient(getBaseURL(t), t)

	if err := client.WaitForAPIToBeHealthy(fiveRetries); err != nil {
		t.Fatal(err)
	}
}
