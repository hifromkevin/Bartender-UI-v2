package database

import (
	"database/sql"
	"fmt"

	"bartender-backend/types"
)

func SeedTools(db *sql.DB) error {
	tools := []types.Tool{
		{
			ID:       "671630a0-3349-461d-a612-75b978660b60",
			ImageURL: "/images/tools/absinthe_spoon.jpg",
			Name:     "Absinthe Spoon",
		},
		{
			ID:       "7e86bd52-9b8c-40a7-ab90-cbb93f504157",
			ImageURL: "/images/tools/bar_spoon.jpg",
			Name:     "Bar Spoon",
		},
		{
			ID:       "33c69a2a-e757-4122-8a53-cb2c8b787e8c",
			ImageURL: "/images/tools/blender.jpg",
			Name:     "Blender",
		},
		{
			ID:       "07c1c81b-a6b7-4987-9f27-bcc1053f1de1",
			ImageURL: "/images/tools/boston_shaker.jpg",
			Name:     "Boston Shaker",
		},
		{
			ID:       "5066fec4-cfd4-426a-942b-bcc29ef71596",
			ImageURL: "/images/tools/bottle_opener.jpg",
			Name:     "Bottle Opener",
		},
		{
			ID:       "b22632aa-3f0c-47bf-8a48-abd5400a438f",
			ImageURL: "/images/tools/carbonation_tool.jpg",
			Name:     "Carbonation Tool",
		},
		{
			ID:       "36d7e434-762e-424a-ad1a-a83c74d96d4f",
			ImageURL: "/images/tools/channel_knife.jpg",
			Name:     "Channel Knife",
		},
		{
			ID:       "7bea95ea-7fea-4efe-a2d4-8b5820f5e3ac",
			ImageURL: "/images/tools/citrus_press.jpg",
			Name:     "Citrus Press",
		},
		{
			ID:       "11eba50c-7cb5-47ff-9588-842e7a9c533f",
			ImageURL: "/images/tools/cobbler_shaker.jpg",
			Name:     "Cobbler Shaker",
		},
		{
			ID:       "2d6a35cd-9ca3-49ea-9b8d-9f03c16485bd",
			ImageURL: "/images/tools/fine_grater.jpg",
			Name:     "Fine Grater",
		},
		{
			ID:       "57c677bd-2025-4edb-b88a-f785e8756217",
			ImageURL: "/images/tools/fine_mesh_sieve.jpg",
			Name:     "Fine Mesh Sieve",
		},
		{
			ID:       "421d1705-027e-48ec-99a6-e04eb270ebdb",
			ImageURL: "/images/tools/hawthorne_strainer.jpg",
			Name:     "Hawthorne Strainer",
		},
		{
			ID:       "aef57b25-d20a-44ff-9554-0dc82cbdcc6c",
			ImageURL: "/images/tools/ice_crusher.jpg",
			Name:     "Ice Crusher",
		},
		{
			ID:       "140dc12c-94a5-406c-a99c-e1fafe22af7e",
			ImageURL: "/images/tools/ice_scoop.jpg",
			Name:     "Ice Scoop",
		},
		{
			ID:       "fb963dc6-46e7-43c2-9be2-7dc222ba31df",
			ImageURL: "/images/tools/jigger.jpg",
			Name:     "Jigger",
		},
		{
			ID:       "9a04473b-ec98-42cd-96d7-a12e52afec8a",
			ImageURL: "/images/tools/julep_strainer.jpg",
			Name:     "Julep Strainer",
		},
		{
			ID:       "6295aa76-76d4-471c-962c-e5117931d977",
			ImageURL: "/images/tools/milk_frother.jpg",
			Name:     "Milk Frother",
		},
		{
			ID:       "816e927f-c0ae-4f50-879c-1012bf418483",
			ImageURL: "/images/tools/mixing_glass.jpg",
			Name:     "Mixing Glass",
		},
		{
			ID:       "f4ad8974-bf4f-4efd-acd4-d3e523953c70",
			ImageURL: "/images/tools/muddler.jpg",
			Name:     "Muddler",
		},
		{
			ID:       "5e786df4-80c9-4922-967a-cc45c2a9ef69",
			ImageURL: "/images/tools/peeler.jpg",
			Name:     "Peeler",
		},
		{
			ID:       "c7ee7b8d-f66b-4de1-91f0-2ce13d5a99bd",
			ImageURL: "/images/tools/shaker.jpg",
			Name:     "Shaker",
		},
		{
			ID:       "df5d86b5-8283-4bfb-9ac2-241a44d0aaf1",
			ImageURL: "/images/tools/smoking_gun.jpg",
			Name:     "Smoking Gun",
		},
		{
			ID:       "f2ec2054-2a91-49e2-aea7-e2f3b73a9cc4",
			ImageURL: "/images/tools/strainer.jpg",
			Name:     "Strainer",
		},
		{
			ID:       "7dbc347b-231a-4493-b0fe-4ca69bbd1d1b",
			ImageURL: "/images/tools/swizzle_stick.jpg",
			Name:     "Swizzle Stick",
		},
		{
			ID:       "88a01c09-8c03-415d-beca-20d3992e5c71",
			ImageURL: "/images/tools/torch.jpg",
			Name:     "Torch",
		},
		{
			ID:       "fa8d8484-6d74-456c-9944-4225bae283c1",
			ImageURL: "/images/tools/wine_key.jpg",
			Name:     "Wine Key",
		},
		{
			ID:       "dede3e98-7f5f-46b6-9f16-5361cfe2284b",
			ImageURL: "/images/tools/zester.jpg",
			Name:     "Zester",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO tools (id, name, image_url) VALUES (?, ?, ?)`)
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

	for _, tool := range tools {
		_, err := stmt.Exec(tool.ID, tool.Name, tool.ImageURL)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
