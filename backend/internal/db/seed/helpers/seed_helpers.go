package helpers

import (
	"database/sql"
	"fmt"
)

func LoadLookupMaps(db *sql.DB, tableNames []string) (map[string]map[string]string, error) {
	lookups := make(map[string]map[string]string)

	for _, table := range tableNames {
		rows, err := db.Query(fmt.Sprintf("SELECT id, name FROM %s", table))
		if err != nil {
			return nil, fmt.Errorf("failed to query table %s: %w", table, err)
		}

		lookup := make(map[string]string)
		for rows.Next() {
			var id, name string
			if err := rows.Scan(&id, &name); err != nil {
				rows.Close()
				return nil, fmt.Errorf("failed to scan row from %s: %w", table, err)
			}
			lookup[name] = id
		}
		rows.Close()

		lookups[table] = lookup
	}

	return lookups, nil
}

func MakeDescriptiveName(parts ...*string) string {
	name := ""
	for _, part := range parts {
		if part != nil {
			if name != "" {
				name += " "
			}
			name += *part
		}
	}
	return name
}

func Pointer(s string) *string {
	return &s
}

func SafeLookup(m map[string]string, key *string, fallback ...string) *string {
	if key == nil {
		return nil
	}

	if val, ok := m[*key]; ok {
		return &val
	}

	if len(fallback) > 0 {
		if val, ok := m[fallback[0]]; ok {
			return &val
		}
		return nil
	}

	return nil
}
