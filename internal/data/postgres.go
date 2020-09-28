package data

import (
	"database/sql"
	"io/ioutil"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", url)
}

// MakeMigration reads and executes models.sql
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
