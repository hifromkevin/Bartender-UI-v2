package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedFlavor(db *sql.DB) error {
	flavorList := []string{
		"Almond",
		"Apple",
		"Apricot",
		"Banana",
		"Birthday Cake",
		"Blackberry",
		"Blueberry",
		"Brown Sugar",
		"Butterscotch",
		"Caramel",
		"Celery",
		"Cherry",
		"Chile",
		"Chocolate",
		"Cinnamon",
		"Coffee",
		"Citrus",
		"Cranberry",
		"Cucumber",
		"Dragon Berry",
		"Dragonfruit",
		"Elderflower",
		"Fig",
		"Grape",
		"Grapefruit",
		"Hazelnut",
		"Irish Cream",
		"Jalapeño",
		"Kiwi",
		"Lemon",
		"Lemonade",
		"Licorice",
		"Lime",
		"Lychee",
		"Mango",
		"Maple",
		"Melon",
		"Orange",
		"Passionfruit",
		"Peach",
		"Peanut Butter",
		"Pear",
		"Pecan",
		"Peppermint",
		"Pineapple",
		"Plum",
		"Pumpkin Spice",
		"Quince",
		"Raspberry",
		"Spiced",
		"Strawberry",
		"Tamarind",
		"Toasted",
		"Tomato",
		"Tropical",
		"Watermelon",
		"Whipped Cream",
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO flavor (id, name) VALUES (?, ?)`)
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

	for _, name := range flavorList {
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
