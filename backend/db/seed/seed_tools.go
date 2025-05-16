package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedTools(db *sql.DB) error {
	tools := []string{
		"muddler",
		"strainer",
		"hawthorne strainer",
		"julep strainer",
		"fine mesh sieve",
		"bar spoon",
		"mixing glass",
		"blender",
		"shaker",
		"boston shaker",
		"cobbler shaker",
		"jigger",
		"zester",
		"peeler",
		"citrus press",
		"channel knife",
		"torch",
		"smoking gun",
		"milk frother",
		"ice crusher",
		"ice scoop",
		"bottle opener",
		"wine key",
		"absinthe spoon",
		"swizzle stick",
		"fine grater",
		"carbonation tool",
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
