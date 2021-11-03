package greet

import (
	"context"
	"fmt"
)

func HelloGreeter(context context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
