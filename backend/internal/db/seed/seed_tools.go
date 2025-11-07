package database

import (
	"database/sql"
	"fmt"
	"log"

	"bartender-backend/internal/db/seed/helpers"
	"bartender-backend/types"
)

func SeedTools(db *sql.DB, atomic bool, jsonPath string) error {
	const insertSQL = `
		INSERT OR IGNORE INTO tools (id, name, image_url)
		VALUES (?, ?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("Prepare tool insert: %w", err)
	}
	defer stmt.Close()

	insertTools := func(exec types.Execer, t types.Tool) error {
		_, err := stmt.Exec(t.ID, t.Name, t.ImageURL)
		return err
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.Tool](
		db,
		jsonPath,
		atomic,
		"tools",
		insertTools,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m🛠️ Tools: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
