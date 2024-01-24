package main

import (
	"jumpstart/database"
	"jumpstart/handlers"
	"jumpstart/models"

	"github.com/labstack/echo/v4"
)

func main() {
	echo := echo.New()
	dbConn := database.GetDbConnection()

	models.CreateUserIfNotExists(dbConn)

	echo.GET("/", handlers.HelloHandler)

	echo.Logger.Fatal(echo.Start(":4000"))
}
