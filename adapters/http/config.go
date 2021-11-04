package http

import (
	"time"

	"go.opentelemetry.io/otel/sdk/trace"
)

type ServerConfig struct {
	Port             string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
	TraceProvider    *trace.TracerProvider
}

func (c ServerConfig) TCPAddress() string {
	return ":" + c.Port
}
