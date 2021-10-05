// +build acceptance

package acceptance

import (
	specifications2 "github.com/quii/go-http-reference-impl/specifications"
	"testing"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	client := NewAPIClient(getBaseURL(t), t)

	if err := client.WaitForAPIToBeHealthy(fiveRetries); err != nil {
		t.Fatal(err)
	}

	specifications2.Greeting(t, client)
	specifications2.RecipeBook(t, client)
}


