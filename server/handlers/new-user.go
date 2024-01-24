package handlers

import (
	"database/sql"
	"jumpstart/models"
	"jumpstart/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegistrationHandler(c echo.Context) error {
	db := c.Get("dbConn").(*sql.DB)

	user := new(models.User)
	user.ID = utils.GenerateUUID()
	if err := c.Bind(user); err != nil {
		log.Println("error binding user:", err)
		return err
	}

	if user.Name == "" || user.Email == "" || user.Role == "" {
		return c.String(http.StatusBadRequest, "missing name, email or role")
	}

	var SqlQuery string = `INSERT INTO users (id, name, email, role) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(SqlQuery, user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		log.Println("error inserting user:", err)
		return err
	}

	return nil
}
