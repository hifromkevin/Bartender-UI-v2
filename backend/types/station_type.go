package types

type Station struct {
  Id				     string						 `json:"id"`
  Position	     int							 `json:"position"`		
  Ingredient     *Ingredient       `json:"ingredient"`
};