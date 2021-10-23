package inmemory

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/quii/go-http-reference-impl/models"
)

type RecipeStore struct {
	store map[string]models.Recipe
}

func NewRecipeStore() *RecipeStore {
	return &RecipeStore{store: make(map[string]models.Recipe)}
}

func (i RecipeStore) GetRecipe(id string) (models.Recipe, error) {
	recipe, exists := i.store[id]

	if !exists {
		return models.Recipe{}, fmt.Errorf("recipe %q does not exist", id)
	}

	return recipe, nil
}

func (i RecipeStore) StoreRecipe(recipe models.Recipe) (string, error) {
	id := uuid.NewString()
	i.store[id] = recipe
	return id, nil
}
