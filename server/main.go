package main

import (
	"dumbsound/config"
	"dumbsound/packages/database"
	"dumbsound/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	database.DatabaseInit()
	config.RunMigration()

	routes.RouteInit(e.Group("dumbsound/api/v1"))

	var PORT = os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + PORT))
}
