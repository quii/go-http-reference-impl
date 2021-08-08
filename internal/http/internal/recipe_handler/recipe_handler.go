package recipe_handler

import (
	"encoding/json"
	"github.com/google/uuid"
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
	id := uuid.NewString()

	_ = rh.storeRecipe(id, recipe) //TODO: handle err

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(RecipeCreateResponse{ID: id})
}

func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = json.NewEncoder(w).Encode(rh.getRecipe(vars))
}

func (rh *RecipeHandler) storeRecipe(id string, recipe RecipeDTO) error {
	return rh.service.StoreRecipe(id, domain.Recipe{
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
