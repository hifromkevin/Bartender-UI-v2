// TODO: Add Pistacchio,
package database

import (
	"database/sql"
	"fmt"
	"log"

	"bartender-backend/internal/db/seed/helpers"
	"bartender-backend/types"
)

func SeedFlavors(db *sql.DB, atomic bool, jsonPath string) error {
	const insertSQL = `
		INSERT OR IGNORE INTO flavors (id, name)
		VALUES (?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("prepare flavor insert: %w", err)
	}
	defer stmt.Close()

	insertFlavors := func(exec types.Execer, t types.Flavor) error {
		_, err := stmt.Exec(t.ID, t.Name)
		return err
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.Flavor](
		db,
		jsonPath,
		atomic,
		"flavors",
		insertFlavors,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m🍦 Flavors: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
