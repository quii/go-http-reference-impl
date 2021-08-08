package recipe_handler

import (
	"github.com/quii/go-http-reference-impl/domain"
)

//go:generate moq -out recipeservice_moq_test.go . RecipeService
type RecipeService interface {
	GetRecipe(id string) (domain.Recipe, error)
	StoreRecipe(id string, recipe domain.Recipe) error
}
