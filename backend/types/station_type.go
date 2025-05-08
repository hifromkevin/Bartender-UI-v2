package types

type Station struct {
  Id				     string						 `json:"id"`
  Position	     int							 `json:"position"`	
  IngredientID	 string						 `json:"ingredient_id"`	
  Ingredient     *Ingredient       `json:"ingredient"`
};