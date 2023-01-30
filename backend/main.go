package main

import (
	"uvent/database"
	"uvent/routers"

	// "github.com/go-playground/validator/v10"
	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connect()
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	routers.SetupRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
