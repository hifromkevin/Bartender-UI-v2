// To run this file, run `go run ./main.go -entity=areaName > fileName.json`
// Remove > fileName.json to print to console instead
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
)

// SeedRow represents a row to be seeded
type SeedRow map[string]string

// generateSeed takes input rows and adds UUIDs
func generateSeed(rows []SeedRow) []SeedRow {
	var out []SeedRow
	for _, row := range rows {
		row["ID"] = uuid.New().String()
		out = append(out, row)
	}
	return out
}

// predefined seed data
var dataSets = map[string][]SeedRow{
	"qualities": {
		{"Name": "Base"},
		{"Name": "Premium"},
		{"Name": "Ultra"},
	},
	"styles": {
		{"Name": "Dark"},
		{"Name": "Light"},
		{"Name": "Silver"},
		{"Name": "Gold"},
		{"Name": "Aged"},
		{"Name": "Overproof"},
	},
	"seasons": {
		{"Name": "Winter"},
		{"Name": "Spring"},
		{"Name": "Summer"},
		{"Name": "Fall"},
	},
}

func main() {
	entity := flag.String("entity", "", "Name of the entity to generate seeds for (e.g., qualities, styles)")
	flag.Parse()

	if *entity == "" {
		fmt.Println("Please specify an entity with -entity flag")
		os.Exit(1)
	}

	rows, ok := dataSets[*entity]
	if !ok {
		fmt.Printf("No dataset found for entity '%s'\n", *entity)
		os.Exit(1)
	}

	generated := generateSeed(rows)

	out, err := json.MarshalIndent(generated, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
