package http

import (
	"github.com/gorilla/mux"

	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal"
	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal/greethandler"
	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal/recipehandler"
	"github.com/quii/go-http-reference-impl/internal/ports"
)

func newRouter(greeter ports.GreeterService, recipeService ports.RecipeService) *mux.Router {
	greetingHandler := greethandler.New(greeter)
	recipeHandler := recipehandler.New(recipeService)

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", greetingHandler.Greet)
	router.HandleFunc("/recipes", recipeHandler.CreateRecipe)
	router.HandleFunc("/recipes/{id}", recipeHandler.GetRecipe)

	return router
}
