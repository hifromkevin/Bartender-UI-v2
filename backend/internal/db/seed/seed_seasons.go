package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedSeason(db *sql.DB) error {
	seasons := []types.Season{
		{
			ID:   "63060729-df9d-42bc-9f23-355fd88afca8",
			Name: "Winter",
		},
		{
			ID:   "72908d6a-9451-4914-b5d6-b16c41c308ba",
			Name: "Spring",
		},
		{
			ID:   "b816a6c0-8ad0-4eee-96a9-89e5268e4420",
			Name: "Summer",
		},
		{
			ID:   "f7db2623-b4c6-4d9c-9c12-4bf034760b9f",
			Name: "Fall",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO seasons (id, name) VALUES (?, ?)`)
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

	for _, season := range seasons {
		_, err := stmt.Exec(season.ID, season.Name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
