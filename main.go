package main

import (
	"goapi2/cmd/handlers"
	"goapi2/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Home)

	e.POST("/users/sign-up", handlers.CreateUser)

	e.POST("/measurements", handlers.CreateMeasurements)

	e.PUT("/users/:id", handlers.HandleUpdateUser)

	e.PUT("/measurements/:id", handlers.HandleUpdateMeasurement)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}}))

	storage.InitDB()

	e.GET("/users/:id", handlers.Profile)

	e.Logger.Fatal(e.Start(":3000"))
}
