package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedFlavorDescriptors(db *sql.DB) error {
	flavorDescriptors := []types.FlavorDescriptor{
		{
			ID:   "60ecc42c-69bc-4116-bf93-b12137de9623",
			Name: "Sweet",
		},
		{
			ID:   "70748bea-73d7-4a93-b963-ab26e6d714e8",
			Name: "Sour",
		},
		{
			ID:   "4e0fcbc4-6c10-48f3-badd-d498f78bd3a5",
			Name: "Bitter",
		},
		{
			ID:   "deda9f5c-e0c2-4e42-823b-ff3fe24ce263",
			Name: "Spicy",
		},
		{
			ID:   "e6df7855-59a9-4678-a590-e3d3fe7bd5b1",
			Name: "Spiced",
		},
		{
			ID:   "c4bd694a-7086-4749-a38a-2662a56ca9b0",
			Name: "Smokey",
		},
		{
			ID:   "475d28cc-dcf5-4d55-b42d-ad592afd7d3c",
			Name: "Toasted",
		},
		{
			ID:   "9e640741-4012-4888-a937-ef4b0fed2c3b",
			Name: "Salted",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO flavor_descriptors (id, name) VALUES (?, ?)`)
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

	for _, flavorDescriptor := range flavorDescriptors {
		_, err := stmt.Exec(flavorDescriptor.ID, flavorDescriptor.Name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
