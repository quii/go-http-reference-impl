package inmemory

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/quii/go-http-reference-impl/application/recipe"
)

type RecipeStore struct {
	store map[string]recipe.Recipe
}

func NewRecipeStore() *RecipeStore {
	return &RecipeStore{store: make(map[string]recipe.Recipe)}
}

func (i RecipeStore) GetRecipe(id string) (recipe.Recipe, error) {
	r, exists := i.store[id]

	if !exists {
		return recipe.Recipe{}, fmt.Errorf("recipe %q does not exist", id)
	}

	return r, nil
}

func (i RecipeStore) StoreRecipe(r recipe.Recipe) (string, error) {
	id := uuid.NewString()
	i.store[id] = r
	return id, nil
}
