//go:build unit
// +build unit

package http_test

import (
	"net/http/httptest"
	"testing"

	http2 "github.com/quii/go-http-reference-impl/adapters/http"
)

func TestNewWebServer(t *testing.T) {
	webServer := http2.NewWebServer(
		http2.ServerConfig{},
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := http2.NewAPIClient(svr.URL, t)

	if err := client.CheckIfHealthy(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
