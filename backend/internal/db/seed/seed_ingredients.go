package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	"bartender-backend/internal/db/seed/helpers" // Go only imports capitalized functions
	"bartender-backend/types"
)

type Execer interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func SeedIngredients(db *sql.DB, atomic bool) error {
	var (
		executor Execer = db
		tx       *sql.Tx
		err      error
	)

	if atomic {
		if tx, err = db.Begin(); err != nil {
			return fmt.Errorf("failed to start transaction: %v", err)
		}
		executor = tx
		log.Println("Transaction started (atomic mode)")
	}

	lookups, err := helpers.LoadLookupMaps(db, []string{
		"IngredientTypes", "Flavors", "Descriptors", "Styles",
		"Brands", "Qualities", "Seasons",
	})
	if err != nil {
		if atomic {
			tx.Rollback()
		}
		return fmt.Errorf("failed to load lookup maps: %v", err)
	}

	ingredientTypes := lookups["IngredientTypes"]
	flavors := lookups["Flavors"]
	descriptors := lookups["Descriptors"]
	styles := lookups["Styles"]
	brands := lookups["Brands"]
	qualities := lookups["Qualities"]
	seasons := lookups["Seasons"]

	ingredients := []types.IngredientSeed{
		{
			IngredientType:   "Liqueur",
			Flavor:           helpers.Pointer("Pear"),
			FlavorDescriptor: helpers.Pointer("Spiced"),
			Style:            helpers.Pointer("Premium"),
			Brand:            helpers.Pointer("St. George"),
			ImageURL:         "/images/stgeorge_spiced_pear.jpg",
			IsAlcoholic:      true,
			IsOrganic:        false,
			IsSeasonal:       false,
			Quality:          helpers.Pointer("Premium"),
			Season:           nil,
		},
		{
			IngredientType:   "Vodka",
			Flavor:           nil,
			FlavorDescriptor: nil,
			Style:            nil,
			Brand:            nil,
			ImageURL:         "/images/vodka.jpg",
			IsAlcoholic:      true,
			IsOrganic:        false,
			IsSeasonal:       false,
			Quality:          helpers.Pointer("Base"),
			Season:           nil,
		},
	}

	rollback := func(reason string, args ...any) error {
		if atomic && tx != nil {
			tx.Rollback()
		}
		return fmt.Errorf(reason, args...)
	}

	for _, i := range ingredients {
		ingredientTypeID, ok := ingredientTypes[i.IngredientType]

		if !ok {
			if atomic {
				tx.Rollback()
			}
			return rollback("unknown ingredient type: %s", i.IngredientType)
		}

		foreignKeys := map[string]*string{
			"flavor_id":            helpers.SafeLookup(flavors, i.Flavor),
			"flavor_descriptor_id": helpers.SafeLookup(descriptors, i.FlavorDescriptor),
			"style_id":             helpers.SafeLookup(styles, i.Style),
			"brand_id":             helpers.SafeLookup(brands, i.Brand),
			"quality_id":           helpers.SafeLookup(qualities, i.Quality, "Base"),
			"season_id":            helpers.SafeLookup(seasons, i.Season),
		}

		descriptiveName := helpers.MakeDescriptiveName(
			i.Brand, i.Style, i.FlavorDescriptor, i.Flavor, helpers.Pointer(i.IngredientType),
		)

		_, loopErr := executor.Exec(`
        INSERT INTO ingredients (
            id,
            ingredient_type_id,
            flavor_id,
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
				ON CONFLICT(descriptive_name) DO NOTHING`,
			uuid.New().String(),
			ingredientTypeID,
			foreignKeys["flavor_id"],
			foreignKeys["flavor_descriptor_id"],
			foreignKeys["style_id"],
			foreignKeys["brand_id"],
			descriptiveName,
			i.ImageURL,
			i.IsAlcoholic,
			i.IsOrganic,
			i.IsSeasonal,
			foreignKeys["quality_id"],
			foreignKeys["season_id"],
		)

		if loopErr != nil {
			if atomic {
				return rollback("insert failed for %s: %v", descriptiveName, loopErr)
			}
			log.Printf("\033[31m❌ Error inserting %s: %v\033[0m", descriptiveName, loopErr)
			continue
		}

		if !atomic {
			log.Printf("\033[32m✅ Inserted ingredient: %s\033[0m", descriptiveName)
		}
	}

	if atomic && tx != nil {
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit failed: %v", err)
		}
		log.Println("Transaction committed successfully")
	}

	return nil
}
