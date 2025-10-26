package database

import (
	"bartender-backend/types"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedIngredientTypes(db *sql.DB) error {
	ingredientTypes := []types.IngredientType{
		{
			ID:   "950a0d19-8b40-4512-a2f3-737a83cd9074",
			Name: "Vodka",
		},
		{
			ID:   "50a8b42d-afc1-4dfd-83dc-69aa4ea695ec",
			Name: "Gin",
		},
		{
			ID:   "6bfb1a75-b0e4-4035-b449-730c2ed1095b",
			Name: "Rum",
		},
		{
			ID:   "96a281cb-638e-4010-aa60-13c59897d593",
			Name: "Tequila",
		},
		{
			ID:   "bafa30b1-bec0-4a65-bcf9-038990e46d3b",
			Name: "Mezcal",
		},
		{
			ID:   "7d5a2001-15bb-4abe-9b9c-cc419d92cbaf",
			Name: "Agave Spirit",
		},
		{
			ID:   "df2090a5-01bb-4cf7-a291-3f5bcdda9e3c",
			Name: "Sake",
		},
		{
			ID:   "5d5a10d4-583e-4ed1-80d6-db22c009afca",
			Name: "Whiskey",
		},
		{
			ID:   "80d1c4d1-b7e3-4053-aea8-15b9a879b81a",
			Name: "Bourbon",
		},
		{
			ID:   "b0a2f62c-1586-4576-80dd-22f4e3b348bd",
			Name: "Scotch",
		},
		{
			ID:   "6cd97d24-7eee-4595-93c5-bea08190aa8f",
			Name: "Everclear",
		},
		{
			ID:   "fb64dc96-cbbf-4606-8b6c-c58886053020",
			Name: "Brandy",
		},
		{
			ID:   "8f8cc281-06bd-4819-897c-522f67be32f4",
			Name: "Applejack",
		},
		{
			ID:   "dab4f2b3-ebed-46c5-aa16-9519e2d14755",
			Name: "Armagnac",
		},
		{
			ID:   "a2ddfd1a-62aa-4db1-8fbc-be350c424786",
			Name: "Cognac",
		},
		{
			ID:   "3f45413d-e56a-48e8-b4a8-8eefa5a2c385",
			Name: "Calvados",
		},
		{
			ID:   "2dc935c3-f3e3-4e26-a74a-f1bfbb4c384f",
			Name: "Grappa",
		},
		{
			ID:   "e0ac60b6-f339-451f-9c93-05e38e825f84",
			Name: "Baijiu",
		},
		{
			ID:   "406f9e1c-9d12-44fa-9198-b0ac267166de",
			Name: "Shōchū",
		},
		{
			ID:   "8071b207-b135-47d3-b92e-5fec93fc9dd3",
			Name: "Schnapps",
		},
		{
			ID:   "77d583c8-bd10-4cd6-8e58-92c2d70372cb",
			Name: "Pisco",
		},
		{
			ID:   "ae92cb08-445f-457b-8eca-f226b6c7c707",
			Name: "Soju",
		},
		{
			ID:   "e2cd41d5-0b18-476d-8603-a1de6cdf3bc9",
			Name: "Cachaça",
		},
		{
			ID:   "38527e32-ab85-446f-9172-11454658287a",
			Name: "Amaretto",
		},
		{
			ID:   "83d0f4a0-a521-45aa-a8cc-548d02f5a9b0",
			Name: "Amaro",
		},
		{
			ID:   "baf253ad-d481-4d6a-8b6d-db25c0a95448",
			Name: "Fernet",
		},
		{
			ID:   "39e2149b-7bea-4e6c-810a-4a2908cca73d",
			Name: "Chartreuse",
		},
		{
			ID:   "42e91d31-3d44-4d36-acdd-35ac57d68c4f",
			Name: "St-Germain",
		},
		{
			ID:   "f47ced13-d92e-490c-8aec-a6b342778840",
			Name: "Campari",
		},
		{
			ID:   "be63d54c-f9b0-4d2f-9126-0857921113c0",
			Name: "Liqueur",
		},
		{
			ID:   "c6904a6c-2346-463d-b093-0f8900e7170f",
			Name: "Triple Sec",
		},
		{
			ID:   "40eb6f2d-d5d4-45c9-8a0f-8d2d5d02f8e2",
			Name: "Curaçao",
		},
		{
			ID:   "17b9bf07-4d69-4d03-86fd-7e4ed622ee4e",
			Name: "Falernum",
		},
		{
			ID:   "86c72d7f-cb22-4d40-ba94-6cfec0e6a18a",
			Name: "Vermouth",
		},
		{
			ID:   "bdfaedb8-7f7a-4b09-b81d-54b296242c8c",
			Name: "Pastis",
		},
		{
			ID:   "320d67c3-6918-4704-acde-1a1cdc4a4c37",
			Name: "Anisette",
		},
		{
			ID:   "dc789ee1-be2f-407c-a40d-028c461b6419",
			Name: "Arak",
		},
		{
			ID:   "9cb8aa8c-13e1-4bad-bb19-3c4ab78a14af",
			Name: "Raki",
		},
		{
			ID:   "cb57e736-c468-481a-bb0d-2f777a6ef679",
			Name: "Absinthe",
		},
		{
			ID:   "33fd8979-b734-4cb7-8064-4bc180159df8",
			Name: "Drambuie",
		},
		{
			ID:   "24676c2b-adf8-4cf7-8ad1-f865d81592bb",
			Name: "Frangelico",
		},
		{
			ID:   "053cabb0-1a8b-410e-a35e-923b3ab1dbca",
			Name: "Galliano",
		},
		{
			ID:   "6a0bd3fd-244a-42cd-9ced-cc0aca5b523f",
			Name: "Juice",
		},
		{
			ID:   "63b8a948-7845-4293-9603-e5e9df1281b2",
			Name: "Tonic Water",
		},
		{
			ID:   "8331bb27-d96b-4250-ba38-6b085572f002",
			Name: "Club Soda",
		},
		{
			ID:   "fd6b1a79-54d0-464e-a239-992e39fd4ee3",
			Name: "Ginger Ale",
		},
		{
			ID:   "7ed53020-0847-4c12-9545-43c765be4823",
			Name: "Ginger Beer",
		},
		{
			ID:   "9c347546-0810-445a-a11e-ad58b8fcca52",
			Name: "Cola",
		},
		{
			ID:   "08f6d148-49fa-43c2-b092-ef45d6365d9c",
			Name: "Soda",
		},
		{
			ID:   "27c3b004-eafe-4a95-9729-b175f68f5df3",
			Name: "Root Beer",
		},
		{
			ID:   "636a3cac-dec7-4740-a719-21c33bc8e993",
			Name: "Energy Drink",
		},
		{
			ID:   "63d7c1bc-6a85-42ac-bc55-2a6628cd3036",
			Name: "Sparkling Water",
		},
		{
			ID:   "ac7d2348-33c6-488a-92a3-df25ffa48010",
			Name: "Syrup",
		},
		{
			ID:   "fa9c0fc0-7082-45f7-9a76-a3863474aa81",
			Name: "Grenadine",
		},
		{
			ID:   "7f7aa611-136c-4e39-8fca-e7eccb721905",
			Name: "Bitters",
		},
		{
			ID:   "2bb5ea84-e5b9-4336-af67-99e3f5748435",
			Name: "Angostura",
		},
		{
			ID:   "72381718-a57c-4c27-9203-a57a1c6b0ef5",
			Name: "Aromatic",
		},
		{
			ID:   "4e293c56-1abc-4561-8d9c-01bcaa75de10",
			Name: "Allspice",
		},
		{
			ID:   "1982bed2-ef34-4765-9cf2-e1a9f1d41334",
			Name: "Cardamom",
		},
		{
			ID:   "9d479170-4a70-452a-94db-1c40d5c0029b",
			Name: "Clove",
		},
		{
			ID:   "bca54684-1cd2-4999-9a88-0dd77b2ce440",
			Name: "Nutmeg",
		},
		{
			ID:   "23868b37-851f-462d-ba9f-3936a14314b4",
			Name: "Saffron",
		},
		{
			ID:   "7ac39403-fedc-45b8-a283-4585f4d57d4a",
			Name: "Star Anise",
		},
		{
			ID:   "e05252d6-6557-431c-9556-cf9a45930f58",
			Name: "Turmeric",
		},
		{
			ID:   "2c83cbd7-daf1-4286-b1b5-a7f7bdbf2c73",
			Name: "Cinnamon Stick",
		},
		{
			ID:   "a224869a-cfe9-410f-abcf-2a868a449fac",
			Name: "Sugar",
		},
		{
			ID:   "5bd500a4-b3d9-4d2e-9626-78f045a3f15c",
			Name: "Salt",
		},
		{
			ID:   "534cf896-7f3f-4e48-ad4e-e75e59c00692",
			Name: "Pepper",
		},
		{
			ID:   "41424d07-88bf-4c1f-a0dd-cd7684055e9c",
			Name: "Tabasco",
		},
		{
			ID:   "52284d54-0503-4db5-90a5-d251807ee9b1",
			Name: "Powder",
		},
		{
			ID:   "1e4288a0-a1a0-40d6-aa67-beadb2781d35",
			Name: "Flakes",
		},
		{
			ID:   "4d7f2a41-7c51-4b35-ab87-c8b4322b8d2a",
			Name: "Oil",
		},
		{
			ID:   "6413f53a-76f7-4146-a202-347152440238",
			Name: "Paste",
		},
		{
			ID:   "5dd4ebe6-26d2-47d2-8798-e045c53972d7",
			Name: "Extract",
		},
		{
			ID:   "09d07af5-6735-4350-bdf9-7130407cedae",
			Name: "Nectar",
		},
		{
			ID:   "88ee29e1-76e0-44c6-8db6-4a5e8f67ce66",
			Name: "Orgeat",
		},
		{
			ID:   "56e1ba0f-9b8f-404b-ae09-eff60aa05cd8",
			Name: "Water",
		},
		{
			ID:   "435f8e73-79e8-43c9-99fb-9c7e06b5853c",
			Name: "Shrub",
		},
		{
			ID:   "1aafaf91-38b0-47f1-9015-b3c412efa367",
			Name: "Rinse",
		},
		{
			ID:   "c4d498e9-f4aa-4a1d-a3d9-f42ac0fd3cd9",
			Name: "Milk",
		},
		{
			ID:   "f485b815-ad96-4edb-aa2e-e4b1b0ed72b2",
			Name: "Cream",
		},
		{
			ID:   "cc1cace5-59ef-4108-8384-7d5aa0d2bdb3",
			Name: "Half and Half",
		},
		{
			ID:   "8406e722-4f75-4c95-9629-bf348e844602",
			Name: "Whipped Cream",
		},
		{
			ID:   "8be433fe-30b4-46f4-978e-7fe802508746",
			Name: "Egg White",
		},
		{
			ID:   "a8e2fb8a-29df-43d9-91b9-3a609ad374a8",
			Name: "Aquafaba",
		},
		{
			ID:   "7860d024-3e20-4b06-bf72-edaed477a19f",
			Name: "Marshmallows",
		},
		{
			ID:   "1060a4ce-6bc3-4e5b-94ff-df814f274d9a",
			Name: "Honeycomb",
		},
		{
			ID:   "c7db7b91-b988-42e2-b559-69ed66f6a537",
			Name: "Cotton Candy",
		},
		{
			ID:   "5c445a65-14eb-4ce7-b35f-479ae062e3fb",
			Name: "Butter",
		},
		{
			ID:   "b1e34e3c-bd20-49d8-ac6c-49764e502829",
			Name: "Spread",
		},
		{
			ID:   "cc85a303-1ac6-4c99-bc3a-b57f8b2e6b3c",
			Name: "Tahini",
		},
		{
			ID:   "e15ba253-82ac-470b-af91-f5ea8dbf2733",
			Name: "Dehydrated Citrus Slice",
		},
		{
			ID:   "e8d0d151-3f1a-4a9e-bbed-17d1543d11f7",
			Name: "Dehydrated Fruit",
		},
		{
			ID:   "eabdfec6-64ad-42c0-be03-7dda6a6d0942",
			Name: "Brandied Cherry",
		},
		{
			ID:   "c510996d-b89c-4006-a090-b50ee5f3c7f1",
			Name: "Maraschino Cherry",
		},
		{
			ID:   "7aadfe9f-8d91-4b66-8775-da6afae4f775",
			Name: "Cherry",
		},
		{
			ID:   "fa73c4bb-2f7e-4b8c-8d17-0539e00a5f3a",
			Name: "Apple",
		},
		{
			ID:   "5b96adbc-12e0-4d94-bf68-d9b61456ec0e",
			Name: "Banana",
		},
		{
			ID:   "827341d5-98cb-4808-aff9-347663f27cc6",
			Name: "Blackberry",
		},
		{
			ID:   "1f084ef0-9302-485a-88d9-08e90e36c9c8",
			Name: "Blueberry",
		},
		{
			ID:   "017332e6-aff8-4b92-a9a2-2daeaaf6f091",
			Name: "Cranberry",
		},
		{
			ID:   "635a0fd8-39ca-4401-8cf0-5a1652150860",
			Name: "Raspberry",
		},
		{
			ID:   "69c3b6e0-5255-4d25-8989-47f7374b55f1",
			Name: "Strawberry",
		},
		{
			ID:   "1203dc98-2745-4551-92c1-2e8bb822cecb",
			Name: "Watermelon",
		},
		{
			ID:   "73c44d55-d7c2-4583-8ea6-ac78a4c10a63",
			Name: "Coconut",
		},
		{
			ID:   "04655494-b593-4775-8b29-6d9294634b0f",
			Name: "Durian",
		},
		{
			ID:   "f7ce25dd-361e-4087-9922-0b67a0e35c95",
			Name: "Guava",
		},
		{
			ID:   "e3b4c08d-e4d0-4e47-a751-4b6d03d0f615",
			Name: "Jackfruit",
		},
		{
			ID:   "55402a1f-5a91-4c3c-afa9-285e1aca50a9",
			Name: "Kumquat",
		},
		{
			ID:   "ff9a5c16-7615-47bb-9f4e-34ea3ecfea81",
			Name: "Papaya",
		},
		{
			ID:   "917de45c-7fd6-4906-a3de-8b3a012fb725",
			Name: "Passionfruit",
		},
		{
			ID:   "524d6a8f-d7be-4b99-b81f-d10d6856ba2b",
			Name: "Pineapple",
		},
		{
			ID:   "bed17e53-c728-4c89-9e17-cff60586b7fa",
			Name: "Tamarind",
		},
		{
			ID:   "dffecf99-ba55-46e3-92cb-2c9a46ae8f0f",
			Name: "Cactus",
		},
		{
			ID:   "a37685ad-e22e-4734-bdb4-ec515f5a6ab5",
			Name: "Dragonfruit",
		},
		{
			ID:   "330a6219-7e2e-4770-a529-8bf72c400c66",
			Name: "Elderberry",
		},
		{
			ID:   "81e5e29c-2977-4bed-bb33-5a00032b8fcb",
			Name: "Gooseberry",
		},
		{
			ID:   "d6305463-3dbd-4d16-abfe-d1ab284e4d15",
			Name: "Kiwi",
		},
		{
			ID:   "d614e886-0e73-4588-8491-c03e09d9c475",
			Name: "Lychee",
		},
		{
			ID:   "ab83c8a6-c0ab-442a-8788-618be613f38c",
			Name: "Mango",
		},
		{
			ID:   "6e601290-a0b8-4e79-a8bd-c3e480e09f47",
			Name: "Mulberry",
		},
		{
			ID:   "03b7e97b-b037-4361-8f13-bfc31022bd8c",
			Name: "Prickly Pear",
		},
		{
			ID:   "63cb6075-9219-42ce-aeca-e9890afbb69e",
			Name: "Yuzu",
		},
		{
			ID:   "d85ccab4-8f2c-4e31-8436-b0ac0a445d64",
			Name: "Slice",
		},
		{
			ID:   "6959aa99-7ce1-4776-9d37-f6025eb03bbd",
			Name: "Persimmon",
		},
		{
			ID:   "e0f95ff1-230a-4628-b812-655098e72318",
			Name: "Peel",
		},
		{
			ID:   "12d808fc-b147-4a65-8884-dff3bbe8fd2e",
			Name: "Zest",
		},
		{
			ID:   "273ff610-72dc-43ee-bf83-763b8529d017",
			Name: "Wheel",
		},
		{
			ID:   "1a2ee2fe-fbca-4671-9b16-30b6c356ace3",
			Name: "Wedge",
		},
		{
			ID:   "67401179-ff35-4cfa-8105-af0af6e8c7e9",
			Name: "Leaf",
		},
		{
			ID:   "3386a236-b646-4129-aa6d-5cf3b97c52f1",
			Name: "Flower",
		},
		{
			ID:   "5c3866c5-c457-43d8-bcbc-b43c38296595",
			Name: "Basil",
		},
		{
			ID:   "061687b0-f818-4056-892c-a4450c1b9314",
			Name: "Thyme",
		},
		{
			ID:   "ad1a21ad-ad10-416c-9684-17575eca6025",
			Name: "Rosemary",
		},
		{
			ID:   "4540acdd-154f-4b54-8301-755d8bd89cc6",
			Name: "Sage",
		},
		{
			ID:   "19018a1f-00bf-4a54-a792-f2578f9ba24e",
			Name: "Lavender",
		},
		{
			ID:   "92d1851c-eeb9-4977-b86a-92d32d21fbd0",
			Name: "Mint",
		},
		{
			ID:   "2b30364b-89e2-4266-97b6-c8f91f0a0d21",
			Name: "Tarragon",
		},
		{
			ID:   "cfa06ece-e667-4a8e-b157-1e788d1f6c63",
			Name: "Celery",
		},
		{
			ID:   "5ebd5ebd-3b58-4691-8581-35eadf8ee078",
			Name: "Carrot",
		},
		{
			ID:   "0496434d-703e-4d81-a7a1-1db770d212d5",
			Name: "Cucumber",
		},
		{
			ID:   "18e3ba88-f2df-428e-a9e8-3aa57ecd7c6c",
			Name: "Olive",
		},
		{
			ID:   "17474d0f-d89d-48bd-856a-b652db12e79b",
			Name: "Pumpkin",
		},
		{
			ID:   "26522fc2-a605-4ff7-9a08-f2e0819d2bae",
			Name: "Potato",
		},
		{
			ID:   "1c35e01f-f122-4569-85e5-1f97d66edd13",
			Name: "Tomato",
		},
		{
			ID:   "804efeba-88f4-48ba-92d6-3af2186d4963",
			Name: "Bacon",
		},
		{
			ID:   "894b92e1-0fd2-4ef9-a934-5629138117b5",
			Name: "Mushroom",
		},
		{
			ID:   "ee28662a-a98f-4b01-a31f-f1cf6238a8c9",
			Name: "Truffle",
		},
		{
			ID:   "ee424883-9e33-4178-9e7b-56192aeb3303",
			Name: "Cashew",
		},
		{
			ID:   "060996ea-261c-445f-8bff-a6a56f1019ce",
			Name: "Macadamia",
		},
		{
			ID:   "105c3df9-fede-4e7b-aa32-b039fe6e9ede",
			Name: "Pistachio",
		},
		{
			ID:   "6b18841d-5f4c-415e-9285-52fd19dbaec9",
			Name: "Walnut",
		},
		{
			ID:   "dbcf95ed-6335-41a0-b5c9-131220f9d39e",
			Name: "Tea",
		},
		{
			ID:   "befefe82-79df-4c4f-a3d1-ec16efbb8540",
			Name: "Ice",
		},
		{
			ID:   "5ed7cee5-0c3a-4955-93f2-98b7d8aa1787",
			Name: "Cocktail Umbrella",
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO ingredient_types (id, name, group_id) VALUES (?, ?, ?)`)
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

	for _, name := range ingredientTypes {
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
