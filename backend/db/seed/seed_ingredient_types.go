package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func SeedIngredientTypes(db *sql.DB) error {
	ingredientTypes := []string{
		"Vodka",
		"Citrus Vodka",
		"Vanilla Vodka",
		"Raspberry Vodka",
		"Peach Vodka",
		"Strawberry Vodka",
		"Blueberry Vodka",
		"Cherry Vodka",
		"Watermelon Vodka",
		"Orange Vodka",
		"Apple Vodka",
		"Banana Vodka",
		"Blackberry Vodka",
		"Kiwi Vodka",
		"Passionfruit Vodka",
		"Dragonfruit Vodka",
		"Lychee Vodka",
		"Melon Vodka",
		"Chocolate Vodka",
		"Cucumber Vodka",
		"Caramel Vodka",
		"Coffee Vodka",
		"Whipped Cream Vodka",
		"Pumpkin Spice Vodka",
		"Peppermint Vodka",
		"Maple Vodka",
		"Spiced Vodka",

		"Gin",

		"Light Rum",
		"Dark Rum",
		"Spiced Rum",
		"Coconut Rum",
		"Overproof Rum",

		"Tequila",
		"Tequila Gold", // 1. What if ingredients are interchangeable? 2. What if they are, and two or more are present on stations?
		"Tequila Silver",
		"Tequila Blanco",
		"Tequila Reposado",
		"Tequila Añejo",

		"Mezcal",

		"Agave Spirit",

		"Pastis",
		"Anisette",
		"Arak",
		"Raki",

		"Grain Alcohol",
		"Everclear",
		"Neutral Grain Spirit",

		"Whiskey",
		"Bourbon",
		"Rye Whiskey",
		"Irish Whiskey",
		"Scotch",

		"Brandy",
		"Applejack",
		"Armagnac",
		"Cognac",
		"Calvados",
		"Grappa",
		"Baijiu",

		"Schnapps",
		"Peppermint Schnapps",
		"Fruit Schnapps",
		"Peach Schnapps",
		"Apple Schnapps",

		"Peach Brandy",
		"Apricot Brandy",
		"Cherry Brandy",
		"Blackberry Brandy",
		"Plum Brandy",
		"Pear Brandy",
		"Quince Brandy",
		"Mirabelle Brandy",
		"Raspberry Brandy",
		"Strawberry Brandy",
		"Blueberry Brandy",
		"Fig Brandy",
		"Fruit Brandy",

		"Pisco",
		"Soju",
		"Cachaça",
		"Absinthe",

		"Amaretto",
		"Amaro",
		"Fernet",
		"Chartreuse",
		"St-Germain",
		"Crème de Violette", // Map to "Cream"
		"Crème de Cassis",
		"Crème de Menthe",
		"Crème de Framboise",
		"Crème de Mûre",
		"Crème de Cerise",
		"Crème de Pêche",
		"Crème de Banane",
		"Crème de Cacao",
		"Crème de Café",
		"Crème de Noisette",
		"Crème de Praliné",
		"Crème de Miel",
		"Crème de Grenadine",
		"Crème de Melon",
		"Crème de Orange",
		"Crème de Citron",
		"Crème de Mandarine",
		"Crème de Pamplemousse",
		"Crème de Yuzu",
		"Crème de Bergamote",
		"Crème de Lavande",
		"Crème de Rose",
		"Crème de Fleur d'Oranger",
		"Crème de Jasmin",
		"Crème de Verveine",

		"Triple Sec",

		"Vermouth",
		"Dry Vermouth",
		"Sweet Vermouth",
		"Sweet Red Vermouth",
		"Campari",

		"Coffee Liqueur",
		"Orange Liqueur",
		"Grand Marnier",
		"Melon Liqueur",
		"Almond Liqueur",
		"Irish Cream Liqueur",
		"Hazelnut Liqueur",
		"Licorice Liqueur",
		"Elderflower Liqueur",
		"Raspberry Liqueur",
		"Strawberry Liqueur",
		"Blueberry Liqueur",
		"Blackberry Liqueur",
		"Cherry Liqueur",
		"Peach Liqueur",
		"Apricot Liqueur",
		"Plum Liqueur",
		"Pear Liqueur",
		"Spiced Pear Liqueur",
		"Fig Liqueur",
		"Quince Liqueur",
		"Passionfruit Liqueur",
		"Lychee Liqueur",
		"Kiwi Liqueur",
		"Dragonfruit Liqueur",
		"Watermelon Liqueur",
		"Banana Liqueur",
		"Caramel Liqueur",
		"Chocolate Liqueur",

		"Blue Curaçao",
		"Falernum",

		"Orange Juice",
		"Cranberry Juice",
		"Pineapple Juice",
		"Lime Juice",
		"Lemon Juice",
		"Grapefruit Juice",
		"Tomato Juice",
		"Apple Juice",
		"Pomegranate Juice",
		"Grape Juice",
		"Mango Juice",
		"Passionfruit Juice",

		"Tonic Water",
		"Club Soda",
		"Ginger Ale",
		"Ginger Beer",

		"Cola",
		"Root Beer",

		"Sparkling Water",
		"Sparkling Lemonade",
		"Sparkling Grapefruit",
		"Sparkling Orange",
		"Sparkling Apple",
		"Sparkling Cranberry",
		"Sparkling Pomegranate",
		"Sparkling Peach",
		"Sparkling Strawberry",
		"Sparkling Raspberry",
		"Sparkling Blueberry",
		"Sparkling Blackberry",
		"Sparkling Kiwi",
		"Sparkling Melon",
		"Sparkling Watermelon",
		"Sparkling Mango",
		"Sparkling Passionfruit",
		"Sparkling Lychee",
		"Sparkling Dragonfruit",
		"Sparkling Coconut",
		"Sparkling Banana",
		"Sparkling Papaya",
		"Sparkling Guava",
		"Sparkling Lemonade",
		"Sparkling Pear",

		"Energy Drink",

		"Lemon-Lime Soda",
		"Root Beer",

		"Sparkling Water",

		"Simple Syrup",
		"Honey Syrup",
		"Agave Syrup",
		"Maple Syrup",
		"Grenadine",
		"Vanilla Syrup",
		"Cinnamon Syrup",
		"Demerara Syrup",
		"Chocolate Syrup",
		"Raspberry Syrup",

		"Angostura Bitters",
		"Orange Bitters",
		"Peychaud's Bitters",
		"Aromatic Bitters",
		"Grapefruit Bitters",
		"Herbal Bitters",
		"Celery Bitters",
		"Chocolate Bitters",
		"Peach Bitters",
		"Cherry Bitters",
		"Cardamom Bitters",
		"Lavender Bitters",
		"Mint Bitters",
		"Black Walnut Bitters",
		"Barrel-Aged Bitters",
		"Bitters",

		"Absinthe Rinse",
		"Agave",
		"Orgeat",
		"Tahini",

		"Tajin",
		"Tabasco",
		"Chili Powder",
		"Chili Flakes",
		"Chili Salt",
		"Chili Oil",
		"Chili Syrup",
		"Chili Paste",

		"Angostura",
		"Orange Curacao",
		"Blue Curacao",
		"Green Curacao",
		"Almond Extract",
		"Almond Oil",

		"Almond Butter",
		"Peanut Butter",

		"Agave Nectar",

		"Almond Syrup",
		"Brown Sugar Syrup",
		"Vanilla Bean Syrup",
		"Caramel Syrup",
		"Chocolate Syrup",
		"Maple Syrup",
		"Peach Syrup",
		"Raspberry Syrup",
		"Strawberry Syrup",
		"Blueberry Syrup",
		"Blackberry Syrup",
		"Cherry Syrup",
		"Orgeat Syrup",

		"Apple Cider Shrub",
		"Balsamic Shrub",

		"Rose Water",
		"Orange Blossom Water",

		"Granulated Sugar",
		"Salt",

		"Milk",
		"Cream",
		"Half and Half",
		"Almond Milk",
		"Oat Milk",

		"Coconut Cream", // Map to Cream of Coconut
		"Egg White",     // Map to Egg Foam
		"Egg Yolk",
		"Aquafaba",

		"Cucumber Slice",

		"Cinnamon Stick",

		"Sugar",
		"Salt",

		"Dehydrated Citrus Slice",

		"Cherry",
		"Maraschino Cherry",
		"Brandied Cherry",

		"Apple",
		"Strawberry",
		"Blueberry",
		"Raspberry",
		"Blackberry",
		"Kiwi",
		"Passionfruit",
		"Dragonfruit",
		"Lychee",
		"Mango",
		"Watermelon",
		"Pineapple Wedge",
		"Peach Slice",
		"Apple Slice",
		"Pear Slice",
		"Orange Peel",
		"Orange Slice",
		"Orange Wheel",
		"Pineapple Slice",
		"Pineapple",

		"Lime Wedge",
		"Lime Wheel",
		"Lemon Wedge",
		"Lemon Twist",

		"Frozen Fruit",
		"Frozen Berries",
		"Frozen Pineapple",
		"Frozen Mango",
		"Frozen Banana",
		"Frozen Coconut",
		"Frozen Yogurt",
		"Frozen Custard",
		"Frozen Sorbet",

		"Basil",
		"Edible Flower",
		"Mint",
		"Pineapple Leaf",

		"Celery",
		"Carrot",
		"Cucumber",
		"Olive",

		"Bacon",

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
