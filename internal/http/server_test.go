// +build unit

package http

import (
	"github.com/quii/hellok8s/acceptance-tests"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := acceptance_tests.NewAPIClient(server.URL)

	acceptance_tests.GreetingAcceptanceTest(t, client)
}