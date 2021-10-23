package greethandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/quii/go-http-reference-impl/internal/ports"
)

type GreetHandler struct {
	greeter ports.GreeterService
}

func New(greeter ports.GreeterService) *GreetHandler {
	return &GreetHandler{greeter: greeter}
}

func (g *GreetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	greeting, err := g.greeter.Greet(mux.Vars(r)["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, greeting)
}
