package recipe

import (
	"fmt"
	"github.com/quii/go-http-reference-impl/domain"
)

type InMemoryRecipeStore struct {
	store map[string]domain.Recipe
}

func NewInMemoryRecipeStore() *InMemoryRecipeStore {
	return &InMemoryRecipeStore{store: make(map[string]domain.Recipe)}
}

func (i InMemoryRecipeStore) GetRecipe(id string) (domain.Recipe, error) {
	recipe, exists := i.store[id]

	if !exists{
		return domain.Recipe{}, fmt.Errorf("recipe %q does not exist", id)
	}

	return recipe, nil
}

func (i InMemoryRecipeStore) StoreRecipe(id string, recipe domain.Recipe) error {
	i.store[id] = recipe
	return nil
}

