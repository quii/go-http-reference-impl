// +build acceptance

package acceptance

import (
	"github.com/quii/go-http-reference-impl/acceptance-criteria"
	"github.com/quii/go-http-reference-impl/acceptance-criteria/adapters"
	"testing"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	client := adapters.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(fiveRetries); err != nil {
		t.Fatal(err)
	}

	acceptance_criteria.GreetingCriteria(t, client)
}


