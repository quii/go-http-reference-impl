// +build acceptance

package acceptance

import (
	"github.com/quii/go-http-reference-impl/internal/acceptance_criteria"
	"github.com/quii/go-http-reference-impl/internal/acceptance_criteria/adapters"
	"testing"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	client := adapters.NewAPIClient(getBaseURL(t), t)

	if err := client.WaitForAPIToBeHealthy(fiveRetries); err != nil {
		t.Fatal(err)
	}

	acceptance_criteria.GreetingCriteria(t, client)
	acceptance_criteria.RecipeStoreCriteria(t, client)
}


