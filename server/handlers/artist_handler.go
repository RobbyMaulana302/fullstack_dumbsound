package handlers

import (
	artistdto "dumbsound/dto/artist"
	resultdto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerArtist struct {
	ArtistRepository repositories.ArtistRepositiry
}

type JsonArtist struct {
	DataArtist interface{} `json:"artist"`
}

func HandlerArtist(ArtistRepository repositories.ArtistRepositiry) *handlerArtist {
	return &handlerArtist{ArtistRepository}
}

func (h *handlerArtist) CreateArtist(c echo.Context) error {
	request := new(artistdto.ArtistRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}

	artistOld, _ := strconv.Atoi(request.Old)
	artist := models.Artist{
		Name:        request.Name,
		Old:         artistOld,
		Type:        request.Type,
		StartCareer: request.StartCareer,
	}

	createArtist, err := h.ArtistRepository.CreateArtist(artist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonArtist{
			DataArtist: convertArtistReponse(createArtist),
		},
	})
}

func (h *handlerArtist) FindArtist(c echo.Context) error {
	artist, err := h.ArtistRepository.FindArtist()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonArtist{
			DataArtist: artist,
		},
	})

}

func (h *handlerArtist) GetArtist(c echo.Context) error {
	artist_id, _ := strconv.Atoi(c.Param("id"))

	artist, err := h.ArtistRepository.GetArtist(artist_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonArtist{
			DataArtist: artist,
		},
	})

}

func (h *handlerArtist) UpdateArtist(c echo.Context) error {
	request := new(artistdto.ArtistUpdateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	artist_id, _ := strconv.Atoi(c.Param("id"))

	artist, err := h.ArtistRepository.GetArtist(artist_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	if request.Name != "" {
		artist.Name = request.Name
	}

	if request.Old != "" {
		old, _ := strconv.Atoi(request.Old)
		artist.Old = old
	}

	if request.Type != "" {
		artist.Type = request.Type
	}

	if request.StartCareer != "" {
		artist.StartCareer = request.StartCareer
	}

	updateArtist, err := h.ArtistRepository.UpdateArtist(artist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonArtist{
			DataArtist: convertArtistReponse(updateArtist),
		},
	})
}

func (h *handlerArtist) DeleteArtist(c echo.Context) error {
	artist_id, _ := strconv.Atoi(c.Param("id"))

	artist, err := h.ArtistRepository.GetArtist(artist_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	deleteArtist, err := h.ArtistRepository.DeleteArtist(artist, artist_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonArtist{
			DataArtist: convertArtistReponse(deleteArtist),
		},
	})
}

func convertArtistReponse(convert models.Artist) artistdto.ArtistResponse {
	return artistdto.ArtistResponse{
		Name:        convert.Name,
		Old:         strconv.Itoa(convert.Old),
		Type:        convert.Type,
		StartCareer: convert.StartCareer,
	}
}
