package http

import (
	"github.com/gorilla/mux"
	"github.com/quii/go-http-reference-impl/internal/greet"
	"github.com/quii/go-http-reference-impl/internal/http/internal"
	"github.com/quii/go-http-reference-impl/internal/http/internal/recipe_handler"
	internal2 "github.com/quii/go-http-reference-impl/internal/recipe"
)

func newRouter() *mux.Router {
	greetingHandler := internal.NewGreetHandler(internal.GreeterFunc(greet.HelloGreeter))
	recipeHandler := recipe_handler.NewRecipeHandler(internal2.NewInMemoryRecipeStore())

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", greetingHandler.Greet)
	router.HandleFunc("/recipes", recipeHandler.CreateRecipe)
	router.HandleFunc("/recipes/{id}", recipeHandler.GetRecipe)

	return router
}
