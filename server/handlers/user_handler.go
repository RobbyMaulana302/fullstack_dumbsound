package handlers

import (
	resultdto "dumbsound/dto/result"
	"dumbsound/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUser(c echo.Context) error {
	User, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data:   User,
	})

}

func (h *handlerUser) GettUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	User, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data:   User,
	})

}
