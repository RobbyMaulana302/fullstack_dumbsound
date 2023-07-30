package handlers

import (
	authdto "dumbsound/dto/auth"
	resultdto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/packages/bcrypt"
	jwtpackage "dumbsound/packages/jwt"
	"dumbsound/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

type JsonAuth struct {
	DataUser interface{} `json:"user"`
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRegister)
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

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	user := models.User{
		Email:     request.Email,
		FullName:  request.FullName,
		Password:  password,
		Gender:    request.Gender,
		Phone:     request.Phone,
		Address:   request.Address,
		ListAs:    false,
		Role:      "user",
		CreatedAd: time.Now(),
		UpdateAd:  time.Now(),
	}

	register, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  "Errro",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonAuth{
			DataUser: register,
		},
	})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.AuthLogin)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	passwordValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !passwordValid {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Status:  "Error",
			Message: "wrong email or password",
		})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["listAs"] = user.ListAs
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, errGenerateToken := jwtpackage.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Email:    user.Email,
		FullName: user.FullName,
		ListAs:   user.ListAs,
		Phone:    user.Phone,
		Address:  user.Address,
		Role:     user.Role,
		Token:    token,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonAuth{
			DataUser: loginResponse,
		},
	})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	userRole := userLogin.(jwt.MapClaims)["role"].(string)

	user, _ := h.AuthRepository.CheckAuth(int(userId), userRole)
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Status: "Success",
		Data: JsonAuth{
			DataUser: user,
		},
	})
}
