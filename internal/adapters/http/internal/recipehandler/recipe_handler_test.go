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

	"github.com/gorilla/mux"
	"github.com/matryer/is"

	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal/recipehandler"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"github.com/quii/go-http-reference-impl/models"
)

func TestGetRecipe(t *testing.T) {
	recipe := models.Recipe{
		Ingredients: []string{"macaroni", "cheese"},
		Directions:  []string{"cook the pasta", "put the cheese"},
		Name:        "Mac and Cheese",
	}

	t.Run("gets recipe by id", func(t *testing.T) {
		is := is.New(t)

		stubService := &ports.RecipeServiceMock{
			GetRecipeFunc: func(id string) (models.Recipe, error) {
				return recipe, nil
			},
		}

		router := mux.NewRouter()
		router.HandleFunc("/recipes/{id}", recipehandler.New(stubService).GetRecipe)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/recipes/%s", "123"), nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)

		var actualRecipe models.Recipe
		is.NoErr(json.Unmarshal(res.Body.Bytes(), &actualRecipe))
		is.Equal(actualRecipe, recipe)
	})

	t.Run("stores recipes", func(t *testing.T) {
		is := is.New(t)

		spyService := &ports.RecipeServiceMock{StoreRecipeFunc: func(recipe models.Recipe) (string, error) {
			return "", nil
		}}

		dto := recipehandler.RecipeDTO{
			Ingredients: recipe.Ingredients,
			Directions:  recipe.Directions,
			Name:        recipe.Name,
		}
		handler := recipehandler.New(spyService)
		req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewReader(dto.ToJSON()))
		res := httptest.NewRecorder()

		handler.CreateRecipe(res, req)

		is.Equal(res.Code, http.StatusCreated)
		is.Equal(len(spyService.StoreRecipeCalls()), 1)
		is.Equal(spyService.StoreRecipeCalls()[0].Recipe, recipe)
	})

	t.Run("it returns a 500 if the recipe call fails", func(t *testing.T) {
		is := is.New(t)

		stubService := &ports.RecipeServiceMock{StoreRecipeFunc: func(recipe models.Recipe) (string, error) {
			return "", errors.New("oh no")
		}}

		handler := recipehandler.New(stubService)
		req := httptest.NewRequest(http.MethodPost, "/recipes", bytes.NewReader(recipehandler.RecipeDTO{}.ToJSON()))
		res := httptest.NewRecorder()

		handler.CreateRecipe(res, req)

		is.Equal(res.Code, http.StatusInternalServerError)
	})
}
