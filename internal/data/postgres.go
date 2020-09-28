package data

import (
	"database/sql"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", url)
}
