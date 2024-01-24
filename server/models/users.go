package models

import "database/sql"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
}

func CreateUserIfNotExists(db *sql.DB) (success bool, error error) {
	SqlQuery := `CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT,
    email TEXT,
    password TEXT,
    role TEXT)`

	_, err := db.Exec(SqlQuery)
	if err != nil {
		return false, err
	}

	return true, nil
}
