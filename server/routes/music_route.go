package routes

import (
	"dumbsound/handlers"
	"dumbsound/packages/database"
	middlewarepackage "dumbsound/packages/middleware"
	"dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func MusicRoutes(e *echo.Group) {
	r := repositories.MakeRepository(database.DB)
	h := handlers.HandlerMusic(r)

	e.POST("/music", middlewarepackage.UploadSong(middlewarepackage.UploadImage(h.CreateMusic)))
	e.GET("/musics", h.FindMusic)
	e.GET("/music/:id", h.GetMusic)

}
