// +build acceptance

package acceptance

import (
	"github.com/quii/go-http-reference-impl/specifications"
	"testing"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	client := NewAPIClient(getBaseURL(t), t)

	if err := client.WaitForAPIToBeHealthy(fiveRetries); err != nil {
		t.Fatal(err)
	}

	specifications.Greeting(t, client)
	specifications.RecipeBook(t, client)
}


