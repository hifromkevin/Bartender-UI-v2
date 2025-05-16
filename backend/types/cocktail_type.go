package types

import "time"

type Cocktail struct {
	ID            string 		`json:"id"`
  Name          string 		`json:"name"`
	Description   string 		`json:"description"`
	IsAlcoholic   bool 		 	`json:"is_alcoholic"`
	ImageURL      string 		`json:"image_url"`
	LastUsedAt		time.Time  	`json:"last_used_at"`
};

type CocktailDetails struct {
	Cocktail     *Cocktail             `json:"cocktail"`
	Steps        []CocktailStep        `json:"steps"`
	Ingredients  []CocktailIngredient  `json:"ingredients"`
	Tools        []CocktailTool        `json:"tools"`
};

type CocktailStep struct {
  ID						string			`json:"id"`
  CocktailID		string			`json:"cocktail_id"`
  StepNumber		int					`json:"step_number"`
  Instruction		string			`json:"instruction"`
};

type CocktailIngredient struct {
	ID							string			`json:"id"`
  CocktailID  		string			`json:"cocktail_id"`
	IngredientID  	string		  `json:"ingredient_id"`
  Amount		    	int					`json:"amount"`
	UnitOfMeasure		string			`json:"unit_of_measure"`
  IsAutomatic			bool				`json:"is_automatic"`
	Category		    string			`json:"category"` // 'mixer' | 'manual_ingredient' | 'garnish';
};

type CocktailTool struct { 
  Id						string			`json:"id"`
	Name					string			`json:"name"`
	Type					string			`json:"type"`
	ImageUrl		  string			`json:"image_url"`
};

type CocktailUsage struct {
	ID         string    `json:"id"`
	CocktailID string    `json:"cocktail_id"`
	MadeAt     time.Time `json:"made_at"`
};