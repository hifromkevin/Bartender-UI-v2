// TODO: Add Pistacchio,
package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"bartender-backend/types"
)

func SeedFlavor(db *sql.DB) error {
	flavorsList := []types.Flavor{
		{
			ID:   "d21cc659-73e3-46e5-91b8-4a6e74b3a37a",
			Name: "Almond",
		},
		{
			ID:   "061eeea2-6dab-4421-b9dd-5d43879e32fd",
			Name: "Apple",
		},
		{
			ID:   "2204ec33-6ef4-4911-9334-0a480116465f",
			Name: "Apricot",
		},
		{
			ID:   "f3e1029a-a58e-4f8c-ba5e-5a3298244e0a",
			Name: "Banana",
		},
		{
			ID:   "d9a90ced-ea77-4871-bf0a-a2c5efdf81ca",
			Name: "Birthday Cake",
		},
		{
			ID:   "5546aae2-cf26-4582-b11d-24e45f514e69",
			Name: "Blackberry",
		},
		{
			ID:   "9e5d9e6a-0ca9-466f-b2fa-f7d36e537e92",
			Name: "Blueberry",
		},
		{
			ID:   "61d78b2e-fed0-4e26-aae6-c745bfff0d93",
			Name: "Brown Sugar",
		},
		{
			ID:   "354421a3-8f78-4e77-9c3d-d4f9d3b319f0",
			Name: "Butterscotch",
		},
		{
			ID:   "1d592049-aead-4562-b17f-00d33eda5ad6",
			Name: "Caramel",
		},
		{
			ID:   "b754b9a6-9b5a-481b-8d9e-d7671aa1c66a",
			Name: "Celery",
		},
		{
			ID:   "0910ace2-3fad-4f81-a89b-1bdd00d98311",
			Name: "Cherry",
		},
		{
			ID:   "151ec488-9da2-4e08-a2e1-b9184fb713b4",
			Name: "Chile",
		},
		{
			ID:   "61c62879-cfcb-4cd7-a4f0-b5feef83abd7",
			Name: "Chocolate",
		},
		{
			ID:   "260161c3-d5cf-4290-a394-8ea136f7dc04",
			Name: "Cinnamon",
		},
		{
			ID:   "d4fe3514-57c2-4ee3-bab4-b5b548d4d89b",
			Name: "Citrus",
		},
		{
			ID:   "e2fdc92e-4b1f-43d7-b498-0a2235f69fd1",
			Name: "Coffee",
		},
		{
			ID:   "6514b991-4091-4706-809f-e81ae161cfff",
			Name: "Cranberry",
		},
		{
			ID:   "8d88a6ce-73fb-4f48-8f58-1ab8e922a97f",
			Name: "Cucumber",
		},
		{
			ID:   "4b44837f-f52b-42cd-b9da-32efad8d92f6",
			Name: "Dragon Berry",
		},
		{
			ID:   "9d55227a-fd58-4892-b0f9-fff7e5a5805c",
			Name: "Dragonfruit",
		},
		{
			ID:   "42f9921f-101b-45eb-bb72-20e70cd99cf5",
			Name: "Elderflower",
		},
		{
			ID:   "864a5317-e93f-499e-9488-c968d440f251",
			Name: "Fig",
		},
		{
			ID:   "fa251ddc-8113-4fbe-b22c-ddabba6ffe0b",
			Name: "Grape",
		},
		{
			ID:   "701d9815-6280-476f-8cb4-64a3df52a610",
			Name: "Grapefruit",
		},
		{
			ID:   "d92ade1e-01f7-435d-9091-4fa0e8f99ac9",
			Name: "Hazelnut",
		},
		{
			ID:   "2995775f-7573-4c38-9fbc-8f88ec152fa6",
			Name: "Irish Cream",
		},
		{
			ID:   "046e3895-ef04-4062-ab78-f2752626f0c4",
			Name: "Jalapeño",
		},
		{
			ID:   "d3a19213-b09b-4707-be53-10110cd5203f",
			Name: "Kiwi",
		},
		{
			ID:   "17aea73b-cc67-499c-8baa-0e10fe855062",
			Name: "Lemon",
		},
		{
			ID:   "a03a657d-8d29-4b1d-9475-67639adc5a4d",
			Name: "Lemonade",
		},
		{
			ID:   "201e66e2-94a1-4bcb-80b7-3398ce733720",
			Name: "Licorice",
		},
		{
			ID:   "6c197e27-bdac-450b-81ce-cf34d76f20ef",
			Name: "Lime",
		},
		{
			ID:   "077540bd-dc79-4c3f-b5f7-9cdc8b668441",
			Name: "Lychee",
		},
		{
			ID:   "5cef7398-70c6-4253-a7cb-610fc516b24c",
			Name: "Mango",
		},
		{
			ID:   "d5e94873-6db0-492c-9693-2aed7dc57ba6",
			Name: "Maple",
		},
		{
			ID:   "d75b264b-275c-4ec8-afbd-a93e13e1de4c",
			Name: "Melon",
		},
		{
			ID:   "1bb04ac5-ef0e-46a0-b49e-8dd9e69a1063",
			Name: "Orange",
		},
		{
			ID:   "fe9a5147-76a7-40c9-ba3e-0283b6abf143",
			Name: "Passionfruit",
		},
		{
			ID:   "74c91c07-5977-4f00-80e3-e4e90a53134d",
			Name: "Peach",
		},
		{
			ID:   "a9ae7d3c-e0aa-4eab-8b6b-8d0b6061d1b5",
			Name: "Peanut Butter",
		},
		{
			ID:   "d93ccd4f-ae75-4719-86e4-6f9f36d4cad4",
			Name: "Pear",
		},
		{
			ID:   "ce289c84-0498-4718-a48e-faf64dfbff28",
			Name: "Pecan",
		},
		{
			ID:   "7c3e1733-0bc2-4ddc-acd0-8e79032032d1",
			Name: "Peppermint",
		},
		{
			ID:   "8625a762-e20a-477a-911c-54e2d7b14832",
			Name: "Pineapple",
		},
		{
			ID:   "a1038dfc-2dbe-43cb-ab9d-cf960c62f640",
			Name: "Plum",
		},
		{
			ID:   "5967771d-47ae-4a5c-bdce-c245bf4ff567",
			Name: "Pumpkin Spice",
		},
		{
			ID:   "643e390b-9155-4df3-984c-68ee9a3ffe8b",
			Name: "Quince",
		},
		{
			ID:   "f8dfe4ff-e776-42eb-9d73-51ad62b68cd6",
			Name: "Raspberry",
		},
		{
			ID:   "9a6cd444-a783-45cb-a147-ac52ba8290a3",
			Name: "Strawberry",
		},
		{
			ID:   "d8a2898c-2c59-4929-a6f3-d9a3df63ae21",
			Name: "Tamarind",
		},
		{
			ID:   "411a4f54-511f-4590-bdd0-662e20d5debe",
			Name: "Toasted",
		},
		{
			ID:   "03bb5c1f-a2c7-43d6-92e6-691e89e69142",
			Name: "Tomato",
		},
		{
			ID:   "55358c6c-2f72-4b5f-a19a-f3fb8627afc3",
			Name: "Tropical",
		},
		{
			ID:   "2012753c-6304-438c-9371-2a78af1d57bd",
			Name: "Watermelon",
		},
		{
			ID:   "3dfba88e-699f-4fa2-9f0a-e324ac644021",
			Name: "Whipped Cream",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO flavors (id, name) VALUES (?, ?)`)
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

	for _, flavor := range flavorsList {
		_, err := stmt.Exec(uuid.New().String(), flavor.Name)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("prepare failed: %v, rollback failed: %v", err, rollbackErr)
			}
			return err
		}
	}

	return tx.Commit()
}
