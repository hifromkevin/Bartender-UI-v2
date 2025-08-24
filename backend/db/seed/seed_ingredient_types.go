package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedIngredientTypes(db *sql.DB) error {
	ingredientTypes := []string{
		"Vodka",
		"Gin",
		"Rum",
		"Light Rum",
		"Dark Rum",
		"Spiced Rum",
		"Coconut Rum",
		"Overproof Rum",
		"Tequila",
		"Tequila Silver",
		"Tequila Gold",
		"Tequila Blanco",
		"Tequila Reposado",
		"Tequila Añejo",
		"Mezcal",
		"Agave Spirit",
		"Sake",
		"Whiskey",
		"Bourbon",
		"Rye Whiskey",
		"Irish Whiskey",
		"Scotch",
		"Everclear",
		"Neutral Grain Spirit",
		"Grain Alcohol",
		"Brandy",
		"Applejack",
		"Armagnac",
		"Cognac",
		"Calvados",
		"Grappa",
		"Baijiu",
		"Shōchū",
		"Schnapps",
		"Pisco",
		"Soju",
		"Cachaça",
		"Amaretto",
		"Amaro",
		"Fernet",
		"Chartreuse",
		"St-Germain",
		"Campari",
		"Liqueur", // "Creme De ..."
		"Triple Sec",
		"Blue Curaçao",
		"Orange Curacao",
		"Green Curacao",
		"Falernum",
		"Dry Vermouth",
		"Sweet Vermouth",
		"Sweet Red Vermouth",
		"Vermouth",

		"Pastis",
		"Anisette",
		"Arak",
		"Raki",
		"Absinthe",

		"Drambuie",
		"Frangelico",
		"Galliano",

		"Juice",
		"Orange Juice", // TODO: These will be used for the juice group with flavor: "Orange", etc. later
		"Apple Juice",
		"Tomato Juice",
		"Carrot Juice",
		"Celery Juice",
		"Beet Juice",
		"Cranberry Juice",
		"Pineapple Juice",
		"Grapefruit Juice",
		"Lemon Juice",
		"Lime Juice",

		"Tonic Water",
		"Club Soda",
		"Ginger Ale",
		"Ginger Beer",
		"Cola",
		"Root Beer",
		"Energy Drink",
		"Lemon-Lime Soda",
		"Sparkling Water",

		"Simple Syrup", // TODO: Move to syrup group
		"Maple Syrup",
		"Agave Syrup",
		"Grenadine",
		"Bitters",
		"Angostura", // TODO: Move to bitters group
		"Aromatic",

		"Allspice",
		"Cardamom",
		"Clove",
		"Nutmeg",
		"Saffron",
		"Star Anise",
		"Turmeric",
		"Cinnamon Stick",
		"Mint Leaf", // TODO: Alias as "Mint"
		"Sugar",

		"Granulated Sugar",
		"Brown Sugar",
		"Powdered Sugar",
		"Salt",
		"Celery Salt",
		"Black Pepper",
		"White Pepper",

		"Tabasco",
		"Chili Powder",
		"Chili Flakes",
		"Chili Salt",
		"Chili Oil",
		"Chili Syrup",
		"Chili Paste",

		"Almond Extract",
		"Almond Oil",

		"Agave Nectar",
		"Orgeat",
		"Rose Water",
		"Orange Blossom Water",
		"Apple Cider Shrub",
		"Balsamic Shrub",
		"Absinthe Rinse",

		"Milk",
		"Cream",
		"Half and Half",
		"Almond Milk",
		"Oat Milk",
		"Coconut Cream",
		"Coconut Milk",
		"Coconut Water",
		"Soy Milk",

		"Egg White",
		"Aquafaba",
		"Marshmallows",
		"Honeycomb",
		"Cotton Candy",

		"Maple Butter",
		"Peanut Butter",
		"Hazelnut Spread",
		"Chocolate Syrup",
		"Tahini",

		"Dehydrated Citrus Slice",
		"Dehydrated Fruit",

		"Brandied Cherry",
		"Maraschino Cherry",
		"Cherry",

		"Apple",
		"Banana",
		"Blackberry",
		"Blueberry",
		"Cranberry",
		"Raspberry",
		"Strawberry",
		"Watermelon",

		"Coconut",
		"Durian",
		"Guava",
		"Jackfruit",
		"Kumquat",
		"Papaya",
		"Passionfruit",
		"Pineapple",
		"Tamarind",

		"Cactus",
		"Dragonfruit",
		"Elderberry",
		"Gooseberry",
		"Kiwi",
		"Lychee",
		"Mango",
		"Mulberry",
		"Passionfruit",
		"Prickly Pear",
		"Yuzu",

		"Apple Slice",
		"Peach Slice",
		"Pear Slice",
		"Orange Slice",
		"Pineapple Slice",
		"Persimmon",
		"Blood Orange",

		"Orange Peel",
		"Orange Zest",
		"Lemon Twist",
		"Lemon Zest",

		"Lime Wedge",
		"Lime Wheel",
		"Lemon Wedge",
		"Orange Wheel",
		"Pineapple Wedge",

		"Edible Flower",
		"Pineapple Leaf",
		"Basil",
		"Thyme",
		"Rosemary",
		"Sage",
		"Lavender",
		"Mint",
		"Tarragon",

		"Celery",
		"Carrot",
		"Cucumber",
		"Cucumber Slice",
		"Olive",
		"Green Olive",
		"Black Olive",
		"Pumpkin",
		"Sweet Potato",
		"Tomato",

		"Bacon",
		"Mushroom",
		"Truffle",

		"Cashew",
		"Macadamia",
		"Pistachio",
		"Walnut",

		"Chai Tea",
		"Earl Grey Tea",
		"Green Tea",

		"Ice",
		"Ice Cubes",
		"Large Ice Cube",
		"Crushed Ice",
		"Shaved Ice",

		"Cocktail Umbrella",
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
