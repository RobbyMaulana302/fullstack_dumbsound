package middlewarepackage

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadImage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileImage, err := c.FormFile("thumbnail")

		if fileImage != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			src, err := fileImage.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()
		}

		c.Set("fileImage", fileImage)
		return next(c)
	}
}

func UploadSong(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileSong, err := c.FormFile("attache")

		if fileSong != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			src, err := fileSong.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()
		}

		c.Set("fileSong", fileSong)
		return next(c)
	}
}
