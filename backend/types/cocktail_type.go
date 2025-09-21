package types

type Cocktail struct {
	ID             string  `db:"id" json:"id"`
	Name           string  `db:"name" json:"name"`
	Description    *string `db:"description" json:"description"`
	IsAlcoholic    bool    `db:"is_alcoholic" json:"is_alcoholic"`
	ImageURL       string  `db:"image_url" json:"image_url"`
	RequiresShaker bool    `db:"requires_shaker" json:"requires_shaker"`
	GlassID        string  `db:"glass_id" json:"glass_id"`
}

type CocktailStep struct {
	ID          string `db:"id" json:"id"`
	CocktailID  string `db:"cocktail_id" json:"cocktail_id"`
	StepNumber  int    `db:"step_number" json:"step_number"`
	Instruction string `db:"instruction" json:"instruction"`
}

type CocktailIngredient struct {
	ID               string  `db:"id" json:"id"`
	CocktailID       string  `db:"cocktail_id" json:"cocktail_id"`
	IngredientID     string  `db:"ingredient_id" json:"ingredient_id"`
	Amount           float64 `db:"amount" json:"amount"`
	UnitOfMeasure    string  `db:"unit_of_measure" json:"unit_of_measure"`
	IsAutomatic      bool    `db:"is_automatic" json:"is_automatic"`
	Category         string  `db:"category" json:"category"`                   // 'mixer' | 'manual_ingredient' | 'garnish';
	HasAutoOption    bool    `db:"has_auto_option" json:"has_auto_option"`     // true if it can be auto OR manual, like "Lime Juice" which can be fresh or bottled
	AlternativeGroup string  `db:"alternative_group" json:"alternative_group"` // same group ID means "or". EX: "limeWheel1234" and "pineChunk5678" both have alternative_group = "garnish1".
}

type CocktailTool struct {
	CocktailID string `db:"cocktail_id" json:"cocktail_id"`
	ToolID     string `db:"tool_id" json:"tool_id"`
}

type CocktailFlavor struct {
	CocktailID string `db:"cocktail_id" json:"cocktail_id"`
	FlavorID   string `db:"flavor_id" json:"flavor_id"`
}

type CocktailFlavorDescriptor struct {
	CocktailID         string `db:"cocktail_id" json:"cocktail_id"`
	FlavorDescriptorID string `db:"flavor_descriptor_id" json:"flavor_descriptor_id"`
}

type CocktailDetails struct {
	Cocktail    *Cocktail            `json:"cocktail"`
	Steps       []CocktailStep       `json:"steps"`
	Ingredients []CocktailIngredient `json:"ingredients"`
	Tools       []CocktailTool       `json:"tools"`
}
