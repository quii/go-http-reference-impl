// +build unit

package http

import (
	"github.com/quii/go-http-reference-impl/acceptance-criteria"
	"github.com/quii/go-http-reference-impl/acceptance-criteria/adapters"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := adapters.NewAPIClient(server.URL, t)

	acceptance_criteria.GreetingCriteria(t, client)
}
