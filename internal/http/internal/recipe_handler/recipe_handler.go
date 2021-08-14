package recipe_handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/quii/go-http-reference-impl/domain"
	"net/http"
)

type RecipeHandler struct {
	service RecipeService
}

func NewRecipeHandler(service RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var recipe RecipeDTO
	_ = json.NewDecoder(r.Body).Decode(&recipe)

	id, _ := rh.storeRecipe(recipe) //TODO: handle err

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(RecipeCreateResponse{ID: id})
}

func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = json.NewEncoder(w).Encode(rh.getRecipe(vars))
}

func (rh *RecipeHandler) storeRecipe(recipe RecipeDTO) (string, error) {
	return rh.service.StoreRecipe(domain.Recipe{
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
