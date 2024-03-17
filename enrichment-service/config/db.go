package config

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
