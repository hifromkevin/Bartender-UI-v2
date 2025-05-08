package types

import "time"

type Cocktail struct {
	ID            string 		`json:"id"`
  Name          string 		`json:"name"`
	Description   string 		`json:"description"`
	IsAlcoholic   bool 		 	`json:"is_alcoholic"`
	ImageURL      string 		`json:"image_url"`
	LastUsedAt		string  	`json:"last_used_at"`
}

type CocktailDetails struct {
	Cocktail     *Cocktail             `json:"cocktail"`
	Steps        []CocktailStep        `json:"steps"`
	Ingredients  []CocktailIngredient  `json:"ingredients"`
	Tools        []CocktailTool        `json:"tools"`
	Glass        *CocktailGlass        `json:"glass"`
}

type CocktailStep struct {
  Id						string			`json:"id"`
  CocktailId		string			`json:"cocktail_id"`
  StepNumber		int					`json:"step_number"`
  Instruction		string			`json:"instruction"`
};

type CocktailIngredient struct {
	Id							string			`json:"id"`
  CocktailId			string			`json:"cocktailId"`
	IngredientId		string		  `json:"ingredientId"`
  Amount		    	int					`json:"amount"`
	UnitOfMeasure		string			`json:"unit_of_measure"`
  IsAutomatic			bool				`json:"is_automatic"`
	Category		    string			`json:"category"` // 'mixer' | 'manual_ingredient' | 'garnish';
};

type CocktailTool struct { 
  Id						string			`json:"id"`
	Name					string			`json:"name"`
	Type					string			`json:"type"`
	ImageUrl		  string			`json:"imageUrl"`
};

type CocktailGlass struct { 
  Id						string			`json:"id"`
	Name					string			`json:"name"`
	Type					string			`json:"type"`
	ImageUrl		  string			`json:"imageUrl"`
};

type CocktailUsage struct {
	ID         string    `json:"id"`
	CocktailID string    `json:"cocktail_id"`
	MadeAt     time.Time `json:"made_at"`
}