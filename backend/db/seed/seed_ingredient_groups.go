package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedIngredientGroups(db *sql.DB) error {
	ingredientGroups := []string{
		"Spirits",
		"Liqueurs",
		"Aperitifs",
		"Digestifs",
		"Bitters",
		"Juices",
		"Sodas",
		"Syrups",
		"Garnishes",
		"Creams/Cremes",
		"Fruits",
		"Vegetables",
		"Herbs and Spices",
		"Sweeteners",
		"Milk",
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
