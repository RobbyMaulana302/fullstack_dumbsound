package handlers

import (
	musicdto "dumbsound/dto/music"
	resultdto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	var err error

	fileImage := c.Get("fileImage").(string)
	fileSong := c.Get("fileSong").(string)

	fmt.Println(fileImage)
	fmt.Println(fileSong)

	artist_id, _ := strconv.Atoi(c.FormValue("artist_id"))

	request := musicdto.MusicRequest{
		Title:     c.FormValue("title"),
		Year:      c.FormValue("year"),
		Thumbnail: fileImage,
		Attache:   fileSong,
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
