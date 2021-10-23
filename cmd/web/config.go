package main

import (
	"time"

	"github.com/quii/go-http-reference-impl/internal/adapters/http"
)

func newDefaultConfig() http.ServerConfig {
	return http.ServerConfig{
		Port:             "8080",
		HTTPReadTimeout:  2 * time.Second,
		HTTPWriteTimeout: 2 * time.Second,
	}
}
