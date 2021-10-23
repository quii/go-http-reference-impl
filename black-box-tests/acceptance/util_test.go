//go:build acceptance
// +build acceptance

package acceptance_test

import (
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
)

const (
	LocalURL = "http://localhost:8080"
)

func getBaseURL(t *testing.T) string {
	url := os.Getenv("BASE_URL")
	if url == "" {
		url = LocalURL
		startWebserver(t)
	}
	return url
}

func startWebserver(t *testing.T) {
	t.Helper()

	compose := testcontainers.NewLocalDockerCompose(
		[]string{"../../docker-compose.yaml"},
		strings.ToLower(uuid.New().String()),
	)
	webContainer := compose.WithCommand([]string{"up", "-d", "web"})
	invokeErr := webContainer.Invoke()

	if invokeErr.Error != nil {
		t.Fatal(invokeErr)
	}

	t.Cleanup(func() {
		compose.Down()
	})
}
