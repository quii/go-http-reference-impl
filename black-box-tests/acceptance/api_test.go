// +build acceptance

package acceptance

import (
	"github.com/matryer/is"
	"github.com/quii/hellok8s"
	"testing"
)

func TestAPI(t *testing.T) {
	client := hello_go_k8s.NewAPIClient(getBaseURL(t))

	if err := client.WaitForAPIToBeHealthy(5); err != nil {
		t.Fatal(err)
	}

	t.Run("it greets with Hello, world!", func(t *testing.T) {
		is := is.New(t)

		greeting, err := client.Greet()
		is.NoErr(err)
		is.Equal(greeting, "Hello, world!")
	})
}



