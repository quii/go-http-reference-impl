// +build acceptance

package acceptance

import (
	"github.com/quii/hellok8s"
	"testing"
)

func TestAPI(t *testing.T) {
	client := hello_go_k8s.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(5); err != nil {
		t.Fatal(err)
	}

	hello_go_k8s.APIAcceptanceTest(t, client)
}


