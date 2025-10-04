package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile("./backend/internal/db/schema/init.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	return db, nil
}
