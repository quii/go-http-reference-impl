package ports

import (
	"github.com/quii/go-http-reference-impl/application/recipe"
)

//go:generate moq -out recipeservice_moq.go . RecipeService
type RecipeService interface {
	GetRecipe(id string) (recipe.Recipe, error)
	StoreRecipe(recipe recipe.Recipe) (id string, err error)
}
