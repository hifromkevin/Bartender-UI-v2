package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedIngredientQuality(db *sql.DB) error {
	ingredientQualities := []types.IngredientQuality{
		{ID: "9469862c-34d1-479f-9af4-0cc0372bd737", Name: "Base"},
		{ID: "9ab2ffc7-05d5-4a8f-b027-173cfcd5b700", Name: "Premium"},
		{ID: "16e6ec8a-3930-460e-b7a7-70d151b41d01", Name: "Ultra"},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO ingredient_quality (id, name) VALUES (?, ?)`)
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

	for _, quality := range ingredientQualities {
		_, err := stmt.Exec(quality.Name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
