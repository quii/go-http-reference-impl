package greet_handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http"
)

type GreetHandler struct {
	greeter ports.GreeterService
}

func NewGreetHandler(greeter ports.GreeterService) *GreetHandler {
	return &GreetHandler{greeter: greeter}
}

func (g *GreetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	greeting, err := g.greeter.Greet(vars["name"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, greeting)
}
