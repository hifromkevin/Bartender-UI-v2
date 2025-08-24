package types

import "time"

type CocktailLog struct {
	ID         string    `json:"id"`
	CocktailID string    `json:"cocktail_id"`
	Timestamp  time.Time `json:"timestamp"`
}
