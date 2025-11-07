package database

import (
	"database/sql"
	"fmt"
	"log"

	"bartender-backend/internal/db/seed/helpers"
	"bartender-backend/types"
)

func SeedFlavorDescriptors(db *sql.DB, atomic bool, jsonPath string) error {
	const insertSQL = `
		INSERT OR IGNORE INTO flavor_descriptors (id, name)
		VALUES (?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("prepare flavor descriptors insert: %w", err)
	}
	defer stmt.Close()

	insertFlavorDescriptors := func(exec types.Execer, t types.FlavorDescriptor) error {
		_, err := stmt.Exec(t.ID, t.Name)
		return err
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.FlavorDescriptor](
		db,
		jsonPath,
		atomic,
		"flavor_descriptors",
		insertFlavorDescriptors,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m➕ Flavor Descriptors: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
