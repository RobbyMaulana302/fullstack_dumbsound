package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	ArtistRoutes(e)
	TransactionRoutes(e)
	MusicRoutes(e)
	UserRoutes(e)
}
