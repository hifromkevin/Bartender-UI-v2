package database

import (
	"database/sql"
	"fmt"
	"log"

	"bartender-backend/internal/db/seed/helpers"
	"bartender-backend/types"
)

func SeedSeasons(db *sql.DB, atomic bool, jsonPath string) error {
	const insertSQL = `
		INSERT OR IGNORE INTO seasons (id, name)
		VALUES (?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("prepare season insert: %w", err)
	}
	defer stmt.Close()

	insertSeasons := func(exec types.Execer, t types.Season) error {
		_, err := stmt.Exec(t.ID, t.Name)
		return err
	}

	inserted, skipped, duration, err := helpers.SeedJSON[types.Season](
		db,
		jsonPath,
		atomic,
		"seasons",
		insertSeasons,
	)

	if err != nil {
		return err
	}

	log.Printf("\033[32m🍁 Seasons: inserted=%d skipped=%d (%s)\033[0m", inserted, skipped, duration)
	return nil
}
