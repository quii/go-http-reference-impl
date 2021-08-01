package main

import (
	"github.com/quii/go-http-reference-impl/internal/http"
	"log"
	"time"
)

func main() {
	server := http.NewWebServer(http.ServerConfig{
		Port:             "8080",
		HTTPReadTimeout:  2 * time.Second,
		HTTPWriteTimeout: 2 * time.Second,
	})

	log.Fatal(server.ListenAndServe())
}
