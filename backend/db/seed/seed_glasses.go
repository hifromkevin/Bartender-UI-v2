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
		{name: "Beer Bottle", image: "/images/glasses/beer_bottle.jpg"},
		{name: "Beer Glass", image: "/images/glasses/beer_glass.jpg"},
		{name: "Beer Mug", image: "/images/glasses/beer_mug.jpg"},
		{name: "Beer Pitcher", image: "/images/glasses/beer_pitcher.jpg"},
		{name: "Beer Stein", image: "/images/glasses/beer_stein.jpg"},
		{name: "Brandy Snifter", image: "/images/glasses/brandy_snifter.jpg"},
		{name: "Caraffe", image: "/images/glasses/caraffe.jpg"},
		{name: "Champagne Bucket", image: "/images/glasses/champagne_bucket.jpg"},
		{name: "Champagne Coupe", image: "/images/glasses/champagne_coupe.jpg"},
		{name: "Champagne Flute", image: "/images/glasses/champagne_flute.jpg"},
		{name: "Champagne Tulip", image: "/images/glasses/champagne_tulip.jpg"},
		{name: "Collins", image: "/images/glasses/collins.jpg"},
		{name: "Coupe", image: "/images/glasses/coupe.jpg"},
		{name: "Copper Mug", image: "/images/glasses/copper_mug.jpg"},
		{name: "Double Old Fashioned", image: "/images/glasses/double_old_fashioned.jpg"},
		{name: "Double Rocks", image: "/images/glasses/double_rocks.jpg"},
		{name: "Diver Glass", image: "/images/glasses/diver_glass.jpg"},
		{name: "Flagon", image: "/images/glasses/flagon.jpg"},
		{name: "Flared", image: "/images/glasses/flared.jpg"},
		{name: "Flared Highball", image: "/images/glasses/flared_highball.jpg"},
		{name: "Flared Rocks", image: "/images/glasses/flared_rocks.jpg"},
		{name: "Flared Shot", image: "/images/glasses/flared_shot.jpg"},
		{name: "Flared Shot Glass", image: "/images/glasses/flared_shot_glass.jpg"},
		{name: "Flute", image: "/images/glasses/flute.jpg"},
		{name: "Highball", image: "/images/glasses/highball.jpg"},
		{name: "Julep", image: "/images/glasses/julep.jpg"},
		{name: "Martini", image: "/images/glasses/martini.jpg"},
		{name: "Mason Jar", image: "/images/glasses/mason_jar.jpg"},
		{name: "Mezcal", image: "/images/glasses/mezcal.jpg"},
		{name: "Mini Coupe", image: "/images/glasses/mini_coupe.jpg"},
		{name: "Nick and Nora", image: "/images/glasses/nick_and_nora.jpg"},
		{name: "Old Fashioned", image: "/images/glasses/old_fashioned.jpg"},
		{name: "Rocks", image: "/images/glasses/rocks.jpg"},
		{name: "Shot", image: "/images/glasses/shot.jpg"},
		{name: "Shot Glass", image: "/images/glasses/shot_glass.jpg"},
		{name: "Sour Glass", image: "/images/glasses/sour_glass.jpg"},
		{name: "Snifter", image: "/images/glasses/snifter.jpg"},
		{name: "Tiki", image: "/images/glasses/tiki.jpg"},
		{name: "Tot", image: "/images/glasses/tot.jpg"},
		{name: "Wine", image: "/images/glasses/wine.jpg"},
		{name: "Pousse Cafe Glass", image: "/images/glasses/pousse_cafe_glass.jpg"},
		{name: "Hurricane Glass", image: "/images/glasses/hurricane_glass.jpg"},
		{name: "Coffee Mug", image: "/images/glasses/coffee_mug.jpg"},
		{name: "Jar", image: "/images/glasses/jar.jpg"},
		{name: "Irish Coffee Cup", image: "/images/glasses/irish_coffee_cup.jpg"},
		{name: "Punch bowl", image: "/images/glasses/punch_bowl.jpg"},
		{name: "Pitcher", image: "/images/glasses/pitcher.jpg"},
		{name: "Pint Glass", image: "/images/glasses/pint_glass.jpg"},
		{name: "Cordial Glass", image: "/images/glasses/cordial_glass.jpg"},
		{name: "Margarita Glass", image: "/images/glasses/margarita_glass.jpg"},
		{name: "Balloon Glass", image: "/images/glasses/balloon.jpg"},
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
