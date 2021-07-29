package http

import "net/http"

func NewWebServer(config ServerConfig) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      newRouter(),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
