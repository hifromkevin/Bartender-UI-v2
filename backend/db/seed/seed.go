package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func SeedData(db *sql.DB) error {
    _, err := db.Exec(`
        INSERT INTO ingredients (
            id, name, type, input_type, is_alcoholic, image_url,
            usage_total_count, usage_last_30_days, usage_last_used_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, uuid.New().String(), "Vodka", "liquid", "automatic", true, "/public/images/smirnoffVodka.jpg", 0, 0, time.Now().Format(time.RFC3339))
    return err
}