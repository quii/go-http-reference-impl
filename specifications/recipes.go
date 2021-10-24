package specifications

import (
	"testing"

	is "github.com/matryer/is"

	"github.com/quii/go-http-reference-impl/application/recipe"
)

type RecipeStoreAdapter interface {
	Save(recipe recipe.Recipe) (id string, err error)
	Get(id string) (recipe.Recipe, error)
}

func RecipeBook(t *testing.T, adapter RecipeStoreAdapter) {
	t.Helper()
	t.Run("it stores recipes and lets you retrieve them", func(t *testing.T) {
		is := is.New(t)

		r := recipe.Recipe{
			Ingredients: []string{"macaroni", "cheese"},
			Directions:  []string{"cook the pasta", "put the cheese"},
			Name:        "Mac and Cheese",
		}

		id, err := adapter.Save(r)
		is.NoErr(err)

		retrievedRecipe, err := adapter.Get(id)
		is.NoErr(err)
		is.Equal(retrievedRecipe, r)
	})
}
