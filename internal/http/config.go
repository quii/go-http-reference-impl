package http

import "time"

type ServerConfig struct {
	Port             string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
}

func (c ServerConfig) TCPAddress() string {
	return ":" + c.Port
}
