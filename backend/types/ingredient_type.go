package types

type Ingredient struct {
	ID                 string  `db:"id" json:"id"`
	IngredientTypeID   string  `db:"ingredient_type_id" json:"ingredient_type_id"`
	FlavorID           *string `db:"flavor_id" json:"flavor_id"`
	FlavorDescriptorID *string `db:"flavor_descriptor_id" json:"flavor_descriptor_id"`
	StyleID            *string `db:"style_id" json:"style_id"`
	BrandID            *string `db:"brand_id" json:"brand_id"`
	DescriptiveName    string  `db:"descriptive_name" json:"descriptive_name"`
	IceType            *string `db:"ice_type" json:"ice_type"`
	ImageURL           string  `db:"image_url" json:"image_url"`
	IsAlcoholic        bool    `db:"is_alcoholic" json:"is_alcoholic"`
	IsOrganic          bool    `db:"is_organic" json:"is_organic"`
	IsSeasonal         bool    `db:"is_seasonal" json:"is_seasonal"`
	QualityID          string  `db:"quality_id" json:"quality_id"` // 'base', 'premium', 'ultra', 'other'
	SeasonID           *string `db:"season_id" json:"season_id"`
}

type IngredientType struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientGroup struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientBrand struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientStyle struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientFlavor struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientFlavorDescriptor struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientQuality struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type IngredientSeed struct {
	IngredientType   string
	Flavor           *string
	FlavorDescriptor *string
	Style            *string
	Brand            *string
	ImageURL         string
	IsAlcoholic      bool
	IsOrganic        bool
	IsSeasonal       bool
	Quality          *string
	Season           *string
}
