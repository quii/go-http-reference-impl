// +build unit

package http

import (
	hellogok8s "github.com/quii/hellok8s"
	"github.com/quii/hellok8s/acceptance-tests"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := hellogok8s.NewAPIClient(server.URL)

	acceptance_tests.GreetingAcceptanceTest(t, client)
}