package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTables creates the necessary tables if they don't exist
func CreateTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS short_urls (
			id VARCHAR(36) PRIMARY KEY,
			url TEXT NOT NULL,
			short_code VARCHAR(10) UNIQUE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			access_count BIGINT NOT NULL DEFAULT 0
		);
		CREATE INDEX IF NOT EXISTS idx_short_code ON short_urls(short_code);
	`

	_, err := db.Exec(query)
	return err
}
