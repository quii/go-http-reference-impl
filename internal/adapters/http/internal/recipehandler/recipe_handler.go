package recipehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/quii/go-http-reference-impl/internal/ports"
	"github.com/quii/go-http-reference-impl/models"
)

type RecipeHandler struct {
	service ports.RecipeService
}

func New(service ports.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var recipe RecipeDTO
	_ = json.NewDecoder(r.Body).Decode(&recipe)

	id, err := rh.storeRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(RecipeCreateResponse{ID: id})
}

func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = json.NewEncoder(w).Encode(rh.getRecipe(vars))
}

func (rh *RecipeHandler) storeRecipe(recipe RecipeDTO) (string, error) {
	return rh.service.StoreRecipe(models.Recipe{
		Ingredients: recipe.Ingredients,
		Directions:  recipe.Directions,
		Name:        recipe.Name,
	})
}

func (rh *RecipeHandler) getRecipe(vars map[string]string) RecipeDTO {
	recipe, _ := rh.service.GetRecipe(vars["id"])
	return RecipeDTO{
		Ingredients: recipe.Ingredients,
		Directions:  recipe.Directions,
		Name:        recipe.Name,
	}
}
