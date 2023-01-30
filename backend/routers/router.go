package routers

import (
	"uvent/routers/api"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	auth := app.Group("/auth")
	auth.POST("/signup", api.Signup)
	auth.POST("/login", api.Login)
	auth.GET("/user", api.GetUserInfo)

	event := app.Group("/event")
	event.POST("/", api.CreateEvent)
	event.GET("/:id", api.GetEvent)
	event.POST("/apply", api.ApplyToEvent)
	event.GET("/latest", api.GetLatestEvent)
}
