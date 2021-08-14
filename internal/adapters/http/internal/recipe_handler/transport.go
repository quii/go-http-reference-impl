package recipe_handler

import "encoding/json"

type RecipeCreateResponse struct {
	ID string `json:"id"`
}

type RecipeDTO struct {
	Ingredients []string `json:"ingredients"`
	Directions  []string `json:"directions"`
	Name        string   `json:"name"`
}

func (r RecipeDTO) ToJSON() []byte {
	out, _ := json.Marshal(r)
	return out
}
