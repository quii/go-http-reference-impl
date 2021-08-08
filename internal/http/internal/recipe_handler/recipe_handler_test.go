package recipe_handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"github.com/quii/go-http-reference-impl/domain"
	"github.com/quii/go-http-reference-impl/internal/http/internal/recipe_handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecipe(t *testing.T) {

	recipe := domain.Recipe{
		Ingredients: []string{"macaroni", "cheese"},
		Directions:  []string{"cook the pasta", "put the cheese"},
		Name:        "Mac and Cheese",
	}

	t.Run("gets recipe by id", func(t *testing.T) {
		is := is.New(t)

		stubService := &recipe_handler.RecipeServiceMock{
			GetRecipeFunc: func(id string) (domain.Recipe, error) {
				return recipe, nil
			},
		}

		router := mux.NewRouter()
		router.HandleFunc("/recipes/{id}", recipe_handler.NewRecipeHandler(stubService).GetRecipe)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/recipes/%s", "123"), nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)

		var actualRecipe domain.Recipe
		is.NoErr(json.Unmarshal(res.Body.Bytes(), &actualRecipe))
		is.Equal(actualRecipe, recipe)

	})

	t.Run("stores recipes", func(t *testing.T) {
		is := is.New(t)

		spyService := &recipe_handler.RecipeServiceMock{StoreRecipeFunc: func(id string, recipe domain.Recipe) error {
			return nil
		}}

		dto := recipe_handler.RecipeDTO{
			Ingredients: recipe.Ingredients,
			Directions:  recipe.Directions,
			Name:        recipe.Name,
		}
		handler := recipe_handler.NewRecipeHandler(spyService)
		req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewReader(dto.ToJSON()))
		res := httptest.NewRecorder()

		handler.CreateRecipe(res, req)

		is.Equal(res.Code, http.StatusCreated)
		is.Equal(len(spyService.StoreRecipeCalls()), 1)
		is.Equal(spyService.StoreRecipeCalls()[0].Recipe, recipe)
	})
}
