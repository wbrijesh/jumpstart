package middleware

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func ContextDB(dbConn *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dbConn", dbConn)
			return next(c)
		}
	}
}
