package main

import (
	"jumpstart/database"
	"jumpstart/handlers"
	customMiddleware "jumpstart/middleware"
	"jumpstart/models"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	dbConn := database.GetDbConnection()

	e.Use(customMiddleware.ContextDB(dbConn))
	e.Use(echoMiddleware.CORS())

	models.CreateUserIfNotExists(dbConn)

	e.GET("/", handlers.HelloHandler)
	e.POST("/api/v0/new-user", handlers.UserRegistrationHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
