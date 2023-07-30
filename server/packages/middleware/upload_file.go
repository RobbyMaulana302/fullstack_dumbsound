package middlewarepackage

import (
	"io"
	"io/ioutil"
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

			tempFile, err := ioutil.TempFile("uploads", "image-*.png")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()

			c.Set("fileImage", data)
			return next(c)
		}

		c.Set("fileImage", "")
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

			tempFile, err := ioutil.TempFile("uploads", "song-*.mp3")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()

			c.Set("fileSong", data)
			return next(c)
		}

		c.Set("fileSong", "")
		return next(c)
	}
}
