package greet

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
)

func HelloGreeter(ctx context.Context, name string) (string, error) {
	_, span := otel.Tracer(name).Start(ctx, "generate-greeting")
	defer span.End()

	return fmt.Sprintf("Hello, %s!", name), nil
}
