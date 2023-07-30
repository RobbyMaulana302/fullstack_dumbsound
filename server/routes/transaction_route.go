package routes

import (
	"dumbsound/handlers"
	"dumbsound/packages/database"
	middlewarepackage "dumbsound/packages/middleware"
	"dumbsound/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	r := repositories.MakeRepository(database.DB)
	h := handlers.HandlerTransaction(r)

	e.POST("/transaction", middlewarepackage.Auth(h.CreateTransaction))
	e.GET("/transactions", middlewarepackage.Auth(h.FindTransaction))
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transaction", middlewarepackage.Auth(h.GetUserTransaction))
	e.DELETE("/transaction/:id", middlewarepackage.Auth(h.DeleteTransaction))
	e.GET("/getpayment/:id", h.GetPayment)
	e.POST("/notification", h.Notification)
}
