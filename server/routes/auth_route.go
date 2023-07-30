package routes

import (
	"dumbsound/handlers"
	"dumbsound/packages/database"
	middlewarepackage "dumbsound/packages/middleware"
	"dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	r := repositories.MakeRepository(database.DB)
	h := handlers.HandlerAuth(r)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middlewarepackage.Auth(h.CheckAuth))
}
