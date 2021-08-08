package greet

import "fmt"

func HelloGreeter(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
