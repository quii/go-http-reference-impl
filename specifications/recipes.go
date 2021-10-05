package specifications

import (
	is "github.com/matryer/is"
	"github.com/quii/go-http-reference-impl/models"
	"testing"
)

type RecipeStoreAdapter interface {
	Save(recipe models.Recipe) (id string, err error)
	Get(id string) (models.Recipe, error)
}

func RecipeBook(t *testing.T, adapter RecipeStoreAdapter) {
	t.Helper()
	t.Run("it stores recipes and lets you retrieve them", func(t *testing.T) {
		is := is.New(t)

		recipe := models.Recipe{
			Ingredients: []string{"macaroni", "cheese"},
			Directions:  []string{"cook the pasta", "put the cheese"},
			Name:        "Mac and Cheese",
		}

		id, err := adapter.Save(recipe)
		is.NoErr(err)

		retrievedRecipe, err := adapter.Get(id)
		is.NoErr(err)
		is.Equal(retrievedRecipe, recipe)
	})
}
