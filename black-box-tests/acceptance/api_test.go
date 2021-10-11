//go:build acceptance
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

	t.Run("api can do greetings", func(t *testing.T) {
		t.Parallel()
		specifications.Greeting(t, client)
	})

	t.Run("api can act as a recipe book", func(t *testing.T) {
		t.Parallel()
		specifications.RecipeBook(t, client)
	})
}
