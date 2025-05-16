package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedGlasses(db *sql.DB) error {
	glasses := []struct {
		name  string
		image string
	}{
		{name: "Beer Bottle", image: "/backend/public/images/glasses/beer_bottle.jpg"},
		{name: "Beer Glass", image: "/backend/public/images/glasses/beer_glass.jpg"},
		{name: "Beer Mug", image: "/backend/public/images/glasses/beer_mug.jpg"},
		{name: "Beer Pitcher", image: "/backend/public/images/glasses/beer_pitcher.jpg"},
		{name: "Beer Stein", image: "/backend/public/images/glasses/beer_stein.jpg"},
		{name: "Brandy Snifter", image: "/backend/public/images/glasses/brandy_snifter.jpg"},
		{name: "Caraffe", image: "/backend/public/images/glasses/caraffe.jpg"},
		{name: "Champagne Bucket", image: "/backend/public/images/glasses/champagne_bucket.jpg"},
		{name: "Champagne Coupe", image: "/backend/public/images/glasses/champagne_coupe.jpg"},
		{name: "Champagne Flute", image: "/backend/public/images/glasses/champagne_flute.jpg"},
		{name: "Champagne Tulip", image: "/backend/public/images/glasses/champagne_tulip.jpg"},
		{name: "Collins", image: "/backend/public/images/glasses/collins.jpg"},
		{name: "Coupe", image: "/backend/public/images/glasses/coupe.jpg"},
		{name: "Copper Mug", image: "/backend/public/images/glasses/copper_mug.jpg"},
		{name: "Double Old Fashioned", image: "/backend/public/images/glasses/double_old_fashioned.jpg"},
		{name: "Double Rocks", image: "/backend/public/images/glasses/double_rocks.jpg"},
		{name: "Diver Glass", image: "/backend/public/images/glasses/diver_glass.jpg"},
		{name: "Flagon", image: "/backend/public/images/glasses/flagon.jpg"},
		{name: "Flared", image: "/backend/public/images/glasses/flared.jpg"},
		{name: "Flared Highball", image: "/backend/public/images/glasses/flared_highball.jpg"},
		{name: "Flared Rocks", image: "/backend/public/images/glasses/flared_rocks.jpg"},
		{name: "Flared Shot", image: "/backend/public/images/glasses/flared_shot.jpg"},
		{name: "Flared Shot Glass", image: "/backend/public/images/glasses/flared_shot_glass.jpg"},
		{name: "Flute", image: "/backend/public/images/glasses/flute.jpg"},
		{name: "Highball", image: "/backend/public/images/glasses/highball.jpg"},
		{name: "Julep", image: "/backend/public/images/glasses/julep.jpg"},
		{name: "Martini", image: "/backend/public/images/glasses/martini.jpg"},
		{name: "Mason Jar", image: "/backend/public/images/glasses/mason_jar.jpg"},
		{name: "Mezcal", image: "/backend/public/images/glasses/mezcal.jpg"},
		{name: "Mini Coupe", image: "/backend/public/images/glasses/mini_coupe.jpg"},
		{name: "Nick and Nora", image: "/backend/public/images/glasses/nick_and_nora.jpg"},
		{name: "Old Fashioned", image: "/backend/public/images/glasses/old_fashioned.jpg"},
		{name: "Rocks", image: "/backend/public/images/glasses/rocks.jpg"},
		{name: "Shot", image: "/backend/public/images/glasses/shot.jpg"},
		{name: "Shot Glass", image: "/backend/public/images/glasses/shot_glass.jpg"},
		{name: "Sour Glass", image: "/backend/public/images/glasses/sour_glass.jpg"},
		{name: "Snifter", image: "/backend/public/images/glasses/snifter.jpg"},
		{name: "Tiki", image: "/backend/public/images/glasses/tiki.jpg"},
		{name: "Tot", image: "/backend/public/images/glasses/tot.jpg"},
		{name: "Wine", image: "/backend/public/images/glasses/wine.jpg"},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO glasses (id, name) VALUES (?, ?)`)
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

	for _, name := range glasses {
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
