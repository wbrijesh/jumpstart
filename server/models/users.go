package models

import "database/sql"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func CreateUserIfNotExists(db *sql.DB) (success bool, error error) {
	SqlQuery := `CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT,
    email TEXT,
    role TEXT)`

	_, err := db.Exec(SqlQuery)
	if err != nil {
		return false, err
	}

	return true, nil
}
