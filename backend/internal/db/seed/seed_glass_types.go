package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedGlassTypes(db *sql.DB) error {
	glassTypes := []types.GlassType{
		{
			ID:       "1c5719a9-ee21-4a5a-a3d4-9a8a361bff6c",
			ImageURL: "/images/glass_types/beer_mug.jpg",
			Name:     "Beer Mug",
		},
		{
			ID:       "77d71313-8083-4996-97bb-22a334fdddbe",
			ImageURL: "/images/glass_types/brandy_snifter.jpg",
			Name:     "Brandy Snifter",
		},
		{
			ID:       "5094a574-7000-4d48-8526-8403215d97f5",
			ImageURL: "/images/glass_types/carafe.jpg",
			Name:     "Carafe",
		},
		{
			ID:       "25c37cfa-3900-48f8-a46f-e4ad811c9231",
			ImageURL: "/images/glass_types/champagne_glass.jpg",
			Name:     "Champagne Glass",
		},
		{
			ID:       "8f21549d-784a-4c4d-a411-30deb60f35fb",
			ImageURL: "/images/glass_types/collins.jpg",
			Name:     "Collins",
		},
		{
			ID:       "4b5dc981-6396-4f4f-ad4a-b7f3a29bb2f3",
			ImageURL: "/images/glass_types/coupe.jpg",
			Name:     "Coupe",
		},
		{
			ID:       "97dd3c66-3fd9-4e90-a324-0417c4015058",
			ImageURL: "/images/glass_types/copper_mug.jpg",
			Name:     "Copper Mug",
		},
		{
			ID:       "06f7e9cc-7577-490a-b240-ee57a30d02e6",
			ImageURL: "/images/glass_types/double_old_fashioned.jpg",
			Name:     "Double Old Fashioned",
		},
		{
			ID:       "f711316f-2932-4bcb-b8dd-5d90be51d92d",
			ImageURL: "/images/glass_types/double_rocks.jpg",
			Name:     "Double Rocks",
		},
		{
			ID:       "7e12b7bb-3839-4cf5-8132-645cc1e379ff",
			ImageURL: "/images/glass_types/diver_glass.jpg",
			Name:     "Diver Glass",
		},
		{
			ID:       "98bd961d-dd52-4148-86d9-7a5b096514a0",
			ImageURL: "/images/glass_types/flagon.jpg",
			Name:     "Flagon",
		},
		{
			ID:       "b5757ae2-edf2-424c-919f-453612985069",
			ImageURL: "/images/glass_types/flared.jpg",
			Name:     "Flared",
		},
		{
			ID:       "60f421e1-449e-4c69-84b0-8d5cce4dbd49",
			ImageURL: "/images/glass_types/flared_highball.jpg",
			Name:     "Flared Highball",
		},
		{
			ID:       "9fbd5a7e-5a3f-44b5-9d42-a8d1f3a91565",
			ImageURL: "/images/glass_types/flared_rocks.jpg",
			Name:     "Flared Rocks",
		},
		{
			ID:       "e96bc7da-203d-4add-8284-c888c0599439",
			ImageURL: "/images/glass_types/flared_shot.jpg",
			Name:     "Flared Shot",
		},
		{
			ID:       "26777ad8-fcd4-4c11-a899-133a7fe714b5",
			ImageURL: "/images/glass_types/flared_shot_glass.jpg",
			Name:     "Flared Shot Glass",
		},
		{
			ID:       "324970d9-efc7-463b-aa71-75e20157a621",
			ImageURL: "/images/glass_types/flute.jpg",
			Name:     "Flute",
		},
		{
			ID:       "12efdb06-ef7f-4334-8294-a59ba2c17403",
			ImageURL: "/images/glass_types/highball.jpg",
			Name:     "Highball",
		},
		{
			ID:       "17fa3256-78bb-4d69-8e35-ea2bb0b0b85b",
			ImageURL: "/images/glass_types/julep.jpg",
			Name:     "Julep",
		},
		{
			ID:       "153c39da-cb16-42e3-9c25-3d688b0d63b5",
			ImageURL: "/images/glass_types/martini.jpg",
			Name:     "Martini",
		},
		{
			ID:       "77280308-f5d9-4154-9c8a-86331acc39f6",
			ImageURL: "/images/glass_types/mason_jar.jpg",
			Name:     "Mason Jar",
		},
		{
			ID:       "0122ece7-f3ae-4dff-9b4e-77efc885bc0f",
			ImageURL: "/images/glass_types/mezcal.jpg",
			Name:     "Mezcal",
		},
		{
			ID:       "66f730eb-d2f9-402e-95f4-67fb6e62b45e",
			ImageURL: "/images/glass_types/mini_coupe.jpg",
			Name:     "Mini Coupe",
		},
		{
			ID:       "65aa39d3-64fb-4d94-bac7-2c579eccfb17",
			ImageURL: "/images/glass_types/nick_and_nora.jpg",
			Name:     "Nick and Nora",
		},
		{
			ID:       "72706eba-d2df-47f1-ad59-7bbaa65133e1",
			ImageURL: "/images/glass_types/old_fashioned.jpg",
			Name:     "Old Fashioned",
		},
		{
			ID:       "09c57ebf-9bd2-4127-a72d-71563dd377af",
			ImageURL: "/images/glass_types/rocks.jpg",
			Name:     "Rocks",
		},
		{
			ID:       "434a2380-4091-409a-98b8-1d25a18fda4b",
			ImageURL: "/images/glass_types/shot.jpg",
			Name:     "Shot",
		},
		{
			ID:       "d5167eb0-eb2d-4373-9d69-7cbed1c72ea2",
			ImageURL: "/images/glass_types/shot_glass.jpg",
			Name:     "Shot Glass",
		},
		{
			ID:       "b3cb5f40-6f9d-4b45-8ac0-4c532ab75d65",
			ImageURL: "/images/glass_types/sour_glass.jpg",
			Name:     "Sour Glass",
		},
		{
			ID:       "03d3ee0a-3e46-49b1-ac50-d3b0559db3ad",
			ImageURL: "/images/glass_types/snifter.jpg",
			Name:     "Snifter",
		},
		{
			ID:       "cd70454b-70bc-4bc5-a6f9-9d4f5bf54fb2",
			ImageURL: "/images/glass_types/tiki.jpg",
			Name:     "Tiki",
		},
		{
			ID:       "eba0d3f9-e6dc-4f6b-95c4-125e61767d52",
			ImageURL: "/images/glass_types/tot.jpg",
			Name:     "Tot",
		},
		{
			ID:       "06c06846-36f5-4b63-b8a8-27e25e660388",
			ImageURL: "/images/glass_types/wine.jpg",
			Name:     "Wine",
		},
		{
			ID:       "e5fac75d-32ce-4990-ae63-16e26d8fd2a3",
			ImageURL: "/images/glass_types/pousse_cafe_glass.jpg",
			Name:     "Pousse Cafe Glass",
		},
		{
			ID:       "7c16032a-df72-4e28-98f8-7f7d25017b0c",
			ImageURL: "/images/glass_types/hurricane_glass.jpg",
			Name:     "Hurricane Glass",
		},
		{
			ID:       "a444c3d9-7757-447a-b72b-bf12d4fa6cf9",
			ImageURL: "/images/glass_types/coffee_mug.jpg",
			Name:     "Coffee Mug",
		},
		{
			ID:       "c371d4c3-fddf-437c-a421-69ae6f63df20",
			ImageURL: "/images/glass_types/jar.jpg",
			Name:     "Jar",
		},
		{
			ID:       "e06df800-0c28-4236-8c87-28c3a180ebc1",
			ImageURL: "/images/glass_types/irish_coffee_cup.jpg",
			Name:     "Irish Coffee Cup",
		},
		{
			ID:       "0896372d-2782-4feb-9945-c7b37ef0b139",
			ImageURL: "/images/glass_types/punch_bowl.jpg",
			Name:     "Punch Bowl",
		},
		{
			ID:       "3366ab1c-ccb8-4f37-919e-ca64e87ab9d7",
			ImageURL: "/images/glass_types/pitcher.jpg",
			Name:     "Pitcher",
		},
		{
			ID:       "0a33b624-8598-4690-a7e9-e29f9d3ed6ee",
			ImageURL: "/images/glass_types/pint_glass.jpg",
			Name:     "Pint Glass",
		},
		{
			ID:       "586bb89b-fe50-41c5-b331-4aef634dbe8b",
			ImageURL: "/images/glass_types/cordial_glass.jpg",
			Name:     "Cordial Glass",
		},
		{
			ID:       "84b11170-5872-495f-916f-ad2d6b09e5d3",
			ImageURL: "/images/glass_types/margarita_glass.jpg",
			Name:     "Margarita Glass",
		},
		{
			ID:       "b3285563-a197-4e53-9b77-c8412fe30557",
			ImageURL: "/images/glass_types/balloon_glass.jpg",
			Name:     "Balloon Glass",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO glass_types (ID, name, image_url) VALUES (?, ?, ?)`)
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

	for _, glass := range glassTypes {
		_, err := stmt.Exec(glass.ID, glass.Name, glass.ImageURL)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("insert failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
