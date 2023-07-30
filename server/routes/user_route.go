package routes

import (
	"dumbsound/handlers"
	"dumbsound/packages/database"
	"dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	r := repositories.MakeRepository(database.DB)
	h := handlers.HandlerUser(r)

	e.GET("/users", h.FindUser)
	e.GET("/user/:id", h.GettUser)

}
