log.Println("Starting DB schema initialization...")

files := []string{
	"./backend/db/schema/ingredients/ingredients.sql",
	"./backend/db/schema/ingredients/ingredient_groups.sql",
	"./backend/db/schema/ingredients/ingredient_types.sql",
	"./backend/db/schema/ingredients/ingredient_usage.sql",
	"./backend/db/schema/flavors/flavors.sql",
	"./backend/db/schema/glassTypes/glass_types.sql",
		"./backend/db/schema/stations/stations.sql",
	"./backend/db/schema/cocktails/cocktails.sql",
	"./backend/db/schema/cocktails/cocktail_usage.sql",
	"./backend/db/schema/cocktails/cocktail_ingredients.sql",
	"./backend/db/schema/cocktails/cocktail_steps.sql",
	"./backend/db/schema/cocktails/cocktail_tools.sql",

}

var fullSchema string
for _, file := range files {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fullSchema += string(content) + "\n"
}

_, err = db.Exec(fullSchema)
if err != nil {
	return nil, err
}

log.Println("DB schema initialization complete.")
