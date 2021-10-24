//go:build unit
// +build unit

package recipehandler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quii/go-http-reference-impl/application/ports"

	"github.com/gorilla/mux"
	"github.com/matryer/is"

	recipehandler2 "github.com/quii/go-http-reference-impl/adapters/http/internal/recipehandler"

	"github.com/quii/go-http-reference-impl/application/recipe"
)

func TestGetRecipe(t *testing.T) {
	r := recipe.Recipe{
		Ingredients: []string{"macaroni", "cheese"},
		Directions:  []string{"cook the pasta", "put the cheese"},
		Name:        "Mac and Cheese",
	}

	t.Run("gets recipe by id", func(t *testing.T) {
		is := is.New(t)

		stubService := &ports.RecipeServiceMock{
			GetRecipeFunc: func(id string) (recipe.Recipe, error) {
				return r, nil
			},
		}

		router := mux.NewRouter()
		router.HandleFunc("/recipes/{id}", recipehandler2.New(stubService).GetRecipe)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/recipes/%s", "123"), nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)

		var actualRecipe recipe.Recipe
		is.NoErr(json.Unmarshal(res.Body.Bytes(), &actualRecipe))
		is.Equal(actualRecipe, r)
	})

	t.Run("stores recipes", func(t *testing.T) {
		is := is.New(t)

		spyService := &ports.RecipeServiceMock{StoreRecipeFunc: func(recipe recipe.Recipe) (string, error) {
			return "", nil
		}}

		dto := recipehandler2.RecipeDTO{
			Ingredients: r.Ingredients,
			Directions:  r.Directions,
			Name:        r.Name,
		}
		handler := recipehandler2.New(spyService)
		req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewReader(dto.ToJSON()))
		res := httptest.NewRecorder()

		handler.CreateRecipe(res, req)

		is.Equal(res.Code, http.StatusCreated)
		is.Equal(len(spyService.StoreRecipeCalls()), 1)
		is.Equal(spyService.StoreRecipeCalls()[0].RecipeMoqParam, r)
	})

	t.Run("it returns a 500 if the recipe call fails", func(t *testing.T) {
		is := is.New(t)

		stubService := &ports.RecipeServiceMock{StoreRecipeFunc: func(recipe recipe.Recipe) (string, error) {
			return "", errors.New("oh no")
		}}

		handler := recipehandler2.New(stubService)
		req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewReader(recipehandler2.RecipeDTO{}.ToJSON()))
		res := httptest.NewRecorder()

		handler.CreateRecipe(res, req)

		is.Equal(res.Code, http.StatusInternalServerError)
	})
}
