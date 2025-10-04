package types

type Station struct {
	ID           string      `db:"id" json:"id"`
	Position     int         `db:"position" json:"position"`
	IngredientID string      `db:"ingredient_id" json:"ingredient_id"`
	Ingredient   *Ingredient `db:"ingredient" json:"ingredient"`
}
