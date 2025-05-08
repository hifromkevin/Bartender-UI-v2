package types

import "time"

type Ingredient struct {
	ID              string `json:"id"`
	MixerType       string `json:"mixer_type"`
	Brand           string `json:"brand"`
	DescriptiveName string `json:"descriptive_name"`
	Type            string `json:"type"`
	InputType       string `json:"input_type"`
	IsAlcoholic     bool   `json:"is_alcoholic"`
	ImageURL        string `json:"image_url"`
	UsageCount      int    `json:"usage_count"`
	LastUsedAt      string `json:"last_used_at"`
}

type IngredientUsage struct {
	ID         		  string    `json:"id"`
	IngredientID 		string    `json:"ingredient_id"`
	MadeAt     		  time.Time `json:"made_at"`
}