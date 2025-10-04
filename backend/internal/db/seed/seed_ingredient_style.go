package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedIngredientStyles(db *sql.DB) error {
	ingredientStyles := []types.IngredientStyle{
		{
			ID:   "cb163e4f-5091-4fe5-9e2c-bb83dd21f097",
			Name: "Dark",
		},
		{
			ID:   "33ab4bf9-f1bb-4a22-908f-592b132b91c9",
			Name: "Light",
		},
		{
			ID:   "4df768c3-66e6-440f-bd0a-3eea2d700de5",
			Name: "Silver",
		},
		{
			ID:   "9b9a4d51-bd9b-417c-9fee-6e5dcaeda3ab",
			Name: "Gold",
		},
		{
			ID:   "9a4b7c10-1263-4fd8-aadc-5f45fdf0b216",
			Name: "Aged",
		},
		{
			ID:   "c16dbec9-d6c3-4429-96d1-761bb287a8a2",
			Name: "Overproof",
		},
		{
			ID:   "ca1f6ce7-82fd-4687-85bf-d4161409394c",
			Name: "Rye",
		},
		{
			ID:   "1421aa8c-a8b6-47f0-87d2-21e0db92d649",
			Name: "Irish",
		},
		{
			ID:   "5801d71b-39d3-470b-9284-72dadface3eb",
			Name: "Blue",
		},
		{
			ID:   "adf0c342-7e6d-4392-b8c2-e648d44de66e",
			Name: "Green",
		},
		{
			ID:   "d33f8faa-bb33-4231-97d3-94a971d025cc",
			Name: "Orange",
		},
		{
			ID:   "37f990e4-a9d4-454c-a5e5-5da8697dad98",
			Name: "Dry",
		},
		{
			ID:   "d474f588-fed6-4bae-85da-4ad22f06c08d",
			Name: "Sweet",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO ingredient_styles (id, name) VALUES (?, ?)`)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
		}
		return err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			fmt.Printf("stmt close failed: %v\n", closeErr)
		}
	}()

	for _, ingredientStyle := range ingredientStyles {
		_, err := stmt.Exec(ingredientStyle.ID, ingredientStyle.Name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
