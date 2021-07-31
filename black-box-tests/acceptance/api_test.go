// +build acceptance

package acceptance

import (
	"github.com/quii/hellok8s/acceptance-tests"
	"testing"
)

const five_retries = 5

func TestGreetingApplication(t *testing.T) {
	client := acceptance_tests.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(five_retries); err != nil {
		t.Fatal(err)
	}

	acceptance_tests.GreetingAcceptanceTest(t, client)
}


