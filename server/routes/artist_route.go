package routes

import (
	"dumbsound/handlers"
	"dumbsound/packages/database"
	middlewarepackage "dumbsound/packages/middleware"
	"dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func ArtistRoutes(e *echo.Group) {
	r := repositories.MakeRepository(database.DB)
	h := handlers.HandlerArtist(r)

	e.POST("/artist", middlewarepackage.Auth(h.CreateArtist))
	e.GET("/artists", middlewarepackage.Auth(h.FindArtist))
	e.GET("/artist/:id", middlewarepackage.Auth(h.GetArtist))
	e.PATCH("/artist/:id", middlewarepackage.Auth(h.UpdateArtist))
	e.DELETE("/artist/:id", middlewarepackage.Auth(h.DeleteArtist))

}
