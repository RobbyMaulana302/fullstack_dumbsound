package middlewarepackage

import (
	"context"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadImage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileImage, err := c.FormFile("thumbnail")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := fileImage.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		var ctx = context.Background()
		var CLOUDE_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		cloudinary, _ := cloudinary.NewFromParams(CLOUDE_NAME, API_KEY, API_SECRET)

		responseImage, err := cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "dumbsound"})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		c.Set("fileImage", responseImage.SecureURL)
		return next(c)
	}
}

func UploadSong(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileSong, err := c.FormFile("attache")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := fileSong.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		var ctx = context.Background()
		var CLOUDE_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		cloudinary, _ := cloudinary.NewFromParams(CLOUDE_NAME, API_KEY, API_SECRET)
		responseSong, err := cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "dumbsound"})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		c.Set("fileSong", responseSong.SecureURL)
		return next(c)
	}
}
