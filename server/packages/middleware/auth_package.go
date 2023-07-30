package middlewarepackage

import (
	resultdto "dumbsound/dto/result"
	jwtpackage "dumbsound/packages/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: "Error", Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtpackage.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: "Error", Message: "unauthorized"})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
