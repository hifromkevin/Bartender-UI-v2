package types

import "time"

type CocktailLog struct {
	ID         string    `db:"id" json:"id"`
	CocktailID string    `db:"cocktail_id" json:"cocktail_id"`
	Timestamp  time.Time `db:"timestamp" json:"timestamp"`
}
