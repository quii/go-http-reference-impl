// +build acceptance

package acceptance

import (
	"github.com/quii/hellok8s/acceptance-criteria"
	"github.com/quii/hellok8s/acceptance-criteria/adapters"
	"testing"
)

const five_retries = 5

func TestGreetingApplication(t *testing.T) {
	client := adapters.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(five_retries); err != nil {
		t.Fatal(err)
	}

	acceptance_criteria.GreetingCriteria(t, client)
}


