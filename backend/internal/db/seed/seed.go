package database

import (
	"database/sql"
)

// SeedAll calls all individual seeders in order
func SeedAll(db *sql.DB) error {
	if err := SeedData(db); err != nil {
		return err
	}
	if err := SeedGlasses(db); err != nil {
		return err
	}
	// Add more seeders here as needed
	return nil
}
