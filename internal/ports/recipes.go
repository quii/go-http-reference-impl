package ports

import (
	"github.com/quii/go-http-reference-impl/models"
)

//go:generate moq -out recipeservice_moq.go . RecipeService
type RecipeService interface {
	GetRecipe(id string) (models.Recipe, error)
	StoreRecipe(recipe models.Recipe) (id string, err error)
}
