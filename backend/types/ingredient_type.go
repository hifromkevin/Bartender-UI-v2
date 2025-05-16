package types

import (
	"time"
)

type Ingredient struct {
	ID               string  `json:"id"`
	IngredientTypeID string  `json:"ingredient_type_id"`
	FlavorID         string  `json:"flavor_id"`
	Brand            string  `json:"brand"`
	DescriptiveName  string  `json:"descriptive_name"`
	ImageURL         string  `json:"image_url"`
	IsAlcoholic      bool    `json:"is_alcoholic"`
	IsOrganic        bool    `json:"is_organic"`
	IsSeasonal       bool    `json:"is_seasonal"`
	QualityTier      *int    `json:"quality_tier"`
	Season           *string `json:"season"`
}

type IngredientUsage struct {
	ID               string    `json:"id"`
	IngredientID     string    `json:"ingredient_id"`
	IngredientAmount float64   `json:"ingredient_amount"`
	UnitOfMeasure    string    `json:"unit_of_measure"`
	UsedAt           time.Time `json:"used_at"`
}

type IngredientTypes struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	GroupID string `json:"group_id"`
}

type IngredientGroups struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
