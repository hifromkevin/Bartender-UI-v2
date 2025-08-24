package types

type Cocktail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAlcoholic bool   `json:"is_alcoholic"`
	ImageURL    string `json:"image_url"`
}

type CocktailDetails struct {
	Cocktail    *Cocktail            `json:"cocktail"`
	Steps       []CocktailStep       `json:"steps"`
	Ingredients []CocktailIngredient `json:"ingredients"`
	Tools       []CocktailTool       `json:"tools"`
}

type CocktailStep struct {
	ID          string `json:"id"`
	CocktailID  string `json:"cocktail_id"`
	StepNumber  int    `json:"step_number"`
	Instruction string `json:"instruction"`
}

type CocktailIngredient struct {
	ID               string `json:"id"`
	CocktailID       string `json:"cocktail_id"`
	IngredientID     string `json:"ingredient_id"`
	Amount           int    `json:"amount"`
	UnitOfMeasure    string `json:"unit_of_measure"`
	IsAutomatic      bool   `json:"is_automatic"`
	Category         string `json:"category"`          // 'mixer' | 'manual_ingredient' | 'garnish';
	HasAutoOption    bool   `json:"has_auto_option"`   // true if it can be auto OR manual, like "Lime Juice" which can be fresh or bottled
	AlternativeGroup string `json:"alternative_group"` // same group ID means "or". EX: "limeWheel1234" and "pineChunk5678" both have alternative_group = "garnish1".
}

type CocktailTool struct {
	Id         string `json:"id"`
	CocktailID string `json:"cocktail_id"`
	ToolID     string `json:"tool_id"`
}
