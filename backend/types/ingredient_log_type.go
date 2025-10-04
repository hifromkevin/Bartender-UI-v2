package types

type IngredientLog struct {
	ID            string  `db:"id" json:"id"`
	IngredientID  string  `db:"ingredient_id" json:"ingredient_id"`
	Amount        float64 `db:"amount" json:"amount"`
	UnitOfMeasure string  `db:"unit_of_measure" json:"unit_of_measure"`
}

// TODO: Make UnitOfMeasure an enum type???
