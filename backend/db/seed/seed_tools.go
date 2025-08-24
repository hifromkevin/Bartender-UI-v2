package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedTools(db *sql.DB) error {
	tools := []string{
		"Absinthe Spoon",
		"Bar Spoon",
		"Blender",
		"Boston Shaker",
		"Bottle Opener",
		"Carbonation Tool",
		"Channel Knife",
		"Citrus Press",
		"Cobbler Shaker",
		"Fine Grater",
		"Fine Mesh Sieve",
		"Hawthorne Strainer",
		"Ice Crusher",
		"Ice Scoop",
		"Jigger",
		"Julep Strainer",
		"Milk Frother",
		"Mixing Glass",
		"Muddler",
		"Peeler",
		"Shaker",
		"Smoking Gun",
		"Strainer",
		"Swizzle Stick",
		"Torch",
		"Wine Key",
		"Zester",
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO tools (id, name) VALUES (?, ?)`)
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

	for _, name := range tools {
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
