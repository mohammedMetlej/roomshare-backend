package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := fmt.Sprintf(
		"host=localhost port = 5433 user = postgres password = root dbname = roomshare_db sslmode = disable",
	)

	var err error
	DB, err = sql.Open("posthgres", connStr)
	if err != nil {
		return err
	}

	return DB.Ping()
}
