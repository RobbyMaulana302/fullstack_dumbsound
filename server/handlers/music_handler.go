package handlers

import (
	"context"
	musicdto "dumbsound/dto/music"
	resultdto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = os.Getenv("PATH_FILE")

type handlerMusic struct {
	MusicRepository repositories.MusicRepositiry
}

type JsonMusic struct {
	DataMusic interface{} `json:"music"`
}

func HandlerMusic(MusicRepository repositories.MusicRepositiry) *handlerMusic {
	return &handlerMusic{MusicRepository}
}

func (h *handlerMusic) CreateMusic(c echo.Context) error {
	var ctx = context.Background()
	var CLOUDE_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	var err error

	fileImage := c.Get("fileImage").(string)
	fileSong := c.Get("fileSong").(string)

	fmt.Println(fileImage)
	fmt.Println(fileSong)

	artist_id, _ := strconv.Atoi(c.FormValue("artist_id"))

	cloudinary, _ := cloudinary.NewFromParams(CLOUDE_NAME, API_KEY, API_SECRET)

	responseImage, err := cloudinary.Upload.Upload(ctx, fileImage, uploader.UploadParams{Folder: "dumbsound"})
	if err != nil {
		fmt.Println(err.Error())
	}

	responseSong, err := cloudinary.Upload.Upload(ctx, fileSong, uploader.UploadParams{Folder: "dumbsound"})

	request := musicdto.MusicRequest{
		Title:     c.FormValue("title"),
		Year:      c.FormValue("year"),
		Thumbnail: responseImage.SecureURL,
		Attache:   responseSong.SecureURL,
		ArtistID:  artist_id,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	music := models.Music{
		Title:     request.Title,
		Year:      request.Year,
		Thumbnail: request.Thumbnail,
		Attache:   request.Attache,
		ArtistID:  request.ArtistID,
	}

	createMusic, err := h.MusicRepository.CreateMusic(music)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonMusic{
			DataMusic: createMusic,
		},
	})
}

func (h *handlerMusic) FindMusic(c echo.Context) error {
	music, err := h.MusicRepository.FindMusic()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonMusic{
			DataMusic: music,
		},
	})
}

func (h *handlerMusic) GetMusic(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	music, err := h.MusicRepository.GetMusic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonMusic{
			DataMusic: music,
		},
	})
}
