package in_mem

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/quii/go-http-reference-impl/models"
)

type InMemoryRecipeStore struct {
	store map[string]models.Recipe
}

func NewRecipeStore() *InMemoryRecipeStore {
	return &InMemoryRecipeStore{store: make(map[string]models.Recipe)}
}

func (i InMemoryRecipeStore) GetRecipe(id string) (models.Recipe, error) {
	recipe, exists := i.store[id]

	if !exists{
		return models.Recipe{}, fmt.Errorf("recipe %q does not exist", id)
	}

	return recipe, nil
}

func (i InMemoryRecipeStore) StoreRecipe(recipe models.Recipe) (string, error) {
	id := uuid.NewString()
	i.store[id] = recipe
	return id, nil
}

