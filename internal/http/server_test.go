// +build unit

package http

import (
	hellogok8s "github.com/quii/hellok8s"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := hellogok8s.NewAPIClient(server.URL)

	hellogok8s.GreetingAcceptanceTest(t, client)
}