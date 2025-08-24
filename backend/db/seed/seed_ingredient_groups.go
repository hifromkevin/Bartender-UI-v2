package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedIngredientGroups(db *sql.DB) error {
	ingredientGroups := []string{
		"Spirits",
		"Brandy",
		"Liqueurs",
		"Aperitifs",
		"Digestifs",
		"Vermouths",
		"Bitters",
		"Juices",
		"Soft Drinks",
		"Syrups",
		"Garnishes",
		"Fruits",
		"Vegetables",
		"Herbs",
		"Sweeteners",
		"Dairy",
		"Sparkling",
		"Frozen",
		"Toppings",
		"Garnishes",
		"Spicy",
		"Rims",
		"Other Flavorings", //EX: "Vanilla Extract", "Almond Extract", "Rose Water", "Orange Blossom Water"
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO ingredient_groups (id, name) VALUES (?, ?)`)
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

	for _, name := range ingredientGroups {
		_, err := stmt.Exec(uuid.New().String(), name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
