package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Greeter interface {
	Greet(name string) (greeting string, err error)
}

type GreeterFunc func(string) (string, error)

func (g GreeterFunc) Greet(name string) (greeting string, err error) {
	return g(name)
}

type GreetHandler struct {
	greeter Greeter
}

func NewGreetHandler(greeter Greeter) *GreetHandler {
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
