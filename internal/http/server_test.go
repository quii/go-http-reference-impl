// +build unit

package http

import (
	"github.com/quii/hellok8s/acceptance-criteria"
	"github.com/quii/hellok8s/acceptance-criteria/adapters"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	server := httptest.NewServer(NewWebServer(ServerConfig{}).Handler)
	defer server.Close()

	client := adapters.NewAPIClient(server.URL)

	acceptance_criteria.GreetingCriteria(t, client)
}