package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", vars["name"])
}
