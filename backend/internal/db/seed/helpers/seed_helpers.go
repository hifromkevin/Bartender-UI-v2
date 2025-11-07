package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"bartender-backend/types"
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

func isSQLiteConstraint(err error) bool {
	return err != nil && (err.Error() == "UNIQUE constraint failed" || err.Error() == "constraint failed")
}

func SeedJSON[T any](
	db *sql.DB,
	jsonPath string,
	atomic bool,
	label string,
	insertFn func(exec types.Execer, t T) error,
) (inserted int, skipped int, duration time.Duration, err error) {

	start := time.Now()

	jsonFile, fileErr := os.ReadFile(jsonPath)

	if fileErr != nil {
		return 0, 0, 0, fmt.Errorf("\033[31m❌ read %s JSON: %w \033[0m", label, fileErr)
	}

	var items []T
	if unmarshalErr := json.Unmarshal(jsonFile, &items); unmarshalErr != nil {
		return 0, 0, 0, fmt.Errorf("\033[31m❌ unmarshal %s JSON: %w \033[0m", label, unmarshalErr)
	}

	var exec types.Execer = db
	var tx *sql.Tx

	if atomic {
		tx, err = db.Begin()
		if err != nil {
			return 0, 0, 0, fmt.Errorf("\033[31m❌ begin %s transaction: %w \033[0m", label, err)
		}
		exec = tx
	}

	for _, item := range items {
		if err := insertFn(exec, item); err != nil {
			if isSQLiteConstraint(err) {
				skipped++
				continue
			}
			if atomic {
				tx.Rollback()
			}
			return inserted, skipped, time.Since(start), fmt.Errorf("\033[32m✅ Insert %s: %w \033[0m", label, err)
		}
		inserted++
	}

	if atomic {
		if err := tx.Commit(); err != nil {
			return inserted, skipped, time.Since(start), fmt.Errorf("\033[32m✅ Commit %s: %w \033[0m", label, err)
		}
	}

	return inserted, skipped, time.Since(start), nil
}
