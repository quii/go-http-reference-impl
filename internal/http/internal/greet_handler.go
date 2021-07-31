package internal

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
