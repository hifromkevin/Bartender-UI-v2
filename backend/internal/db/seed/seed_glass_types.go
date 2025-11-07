package database

import (
	"database/sql"
	"fmt"
	"log"

	"bartender-backend/internal/db/seed/helpers"
	"bartender-backend/types"
)

func SeedGlassTypes(db *sql.DB, atomic bool, jsonPath string) error {
	const insertSQL = `
		INSERT OR IGNORE INTO glass_types (id, name, image_url)
		VALUES (?, ?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("prepare glass type insert: %w", err)
	}
	defer stmt.Close()

	insertGlassTypes := func(exec types.Execer, t types.GlassType) error {
		_, err := stmt.Exec(t.ID, t.Name, t.ImageURL)
		return err
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.GlassType](
		db,
		jsonPath,
		atomic,
		"glass_types",
		insertGlassTypes,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m🍸 Glass Types: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
