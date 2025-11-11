package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	"bartender-backend/internal/db/seed/helpers" // Go only imports capitalized functions
	"bartender-backend/types"
)

const insertIngredientSQL = `
    INSERT INTO ingredients (
        id,
        ingredient_type_id,
        flavor_ids,
        flavor_descriptor_id,
        style_id,
        brand_id,
        descriptive_name,
        image_url,
        is_alcoholic,
        is_organic,
        is_seasonal,
        quality_id,
        season_id
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    ON CONFLICT(descriptive_name) DO NOTHING
`

func insertIngredient(
	exec types.Execer,
	item types.IngredientSeed,
	lookups map[string]map[string]string,
) error {
	ingredientTypes := lookups["ingredient_types"]
	descriptors := lookups["flavor_descriptors"]
	styles := lookups["ingredient_styles"]
	brands := lookups["ingredient_brands"]
	qualities := lookups["ingredient_quality"]
	seasons := lookups["seasons"]

	ingredientTypeID, ok := ingredientTypes[item.IngredientType]
	if !ok {
		return fmt.Errorf("\033[31m❌ unknown ingredient type: %s \033[0m", item.IngredientType)
	}

	foreignKeys := map[string]*string{
		"flavor_descriptor_id": helpers.SafeLookup(descriptors, item.FlavorDescriptor),
		"style_id":             helpers.SafeLookup(styles, item.Style),
		"brand_id":             helpers.SafeLookup(brands, item.Brand),
		"quality_id":           helpers.SafeLookup(qualities, item.Quality, "Base"),
		"season_id":            helpers.SafeLookup(seasons, item.Season),
	}

	var flavorIds *string
	if len(item.FlavorIDs) > 0 {
		flavor, err := json.Marshal(item.FlavorIDs)
		if err != nil {
			return fmt.Errorf("failed to marshal flavor IDs: %w", err)
		}
		flavorIds = helpers.Pointer(string(flavor))
	}


	descriptiveName := helpers.MakeDescriptiveName(
		item.Brand,
		item.Style,
		item.FlavorDescriptor,
		item.Flavor,
		helpers.Pointer(item.IngredientType),
	)

	_, err := exec.Exec(insertIngredientSQL,
		uuid.New().String(),
		ingredientTypeID,
		flavorIds,
		foreignKeys["flavor_descriptor_id"],
		foreignKeys["style_id"],
		foreignKeys["brand_id"],
		descriptiveName,
		item.ImageURL,
		item.IsAlcoholic,
		item.IsOrganic,
		item.IsSeasonal,
		foreignKeys["quality_id"],
		foreignKeys["season_id"],
	)

	return err
}

func SeedIngredients(db *sql.DB, atomic bool, jsonPath string) error {
	lookups, err := helpers.LoadLookupMaps(db, []string{
		"ingredient_types",
		"ingredient_flavors",
		"flavor_descriptors",
		"ingredient_styles",
		"ingredient_brands",
		"ingredient_quality",
		"seasons",

	})
	if err != nil {
		return fmt.Errorf("\033[31m❌ failed to load lookup maps: %v \033[0m", err)
	}

	insertFn := func(exec types.Execer, item types.IngredientSeed) error {
		return insertIngredient(exec, item, lookups)
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.IngredientSeed](
		db,
		jsonPath,
		atomic,
		"ingredients",
		insertFn,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m🌿 Ingredients: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
