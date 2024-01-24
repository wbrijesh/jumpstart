package database

import (
	"database/sql"
	"fmt"
	"jumpstart/utils"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dbURL string = utils.GetEnv("DB_URL")

func GetDbConnection() *sql.DB {
	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbURL, err)
		os.Exit(1)
	}

	return db
}
