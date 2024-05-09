package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
	"github.com/sawalreverr/bebastukar-be/pkg"
)

type authHandler struct {
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(uc usecase.UserUsecase) AuthHandler {
	return &authHandler{userUsecase: uc}
}

func (h *authHandler) RegisterHandler(c echo.Context) error {
	var userRequest dto.UserCredential

	if err := c.Bind(&userRequest); err != nil {
		bindErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	if err := c.Validate(userRequest); err != nil {
		validateErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, validateErr)
	}

	newUser, err := h.userUsecase.RegisterUser(userRequest)

	if err != nil {
		if errors.Is(err, pkg.ErrDataAlreadyExist) {
			emailExist := helper.ResponseData(http.StatusConflict, "email already registered!", nil)
			return c.JSON(http.StatusConflict, emailExist)
		}

		internalErr := helper.ResponseData(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	responseUser := dto.RegisterResponse{
		ID:          newUser.ID,
		Name:        newUser.Name,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
	}

	response := helper.ResponseData(http.StatusCreated, "Register Successfully!", responseUser)
	return c.JSON(http.StatusCreated, response)
}

func (h *authHandler) LoginHandler(c echo.Context) error {
	var userRequest dto.Login

	if err := c.Bind(&userRequest); err != nil {
		bindErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	if err := c.Validate(&userRequest); err != nil {
		validateErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, validateErr)
	}

	token, err := h.userUsecase.LoginUser(userRequest.Email, userRequest.Password)

	if err != nil {
		if errors.Is(err, pkg.ErrRecordNotFound) {
			wrongCred := helper.ResponseData(http.StatusConflict, "email or password invalid!", nil)
			return c.JSON(http.StatusConflict, wrongCred)
		}

		internalErr := helper.ResponseData(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	responseUser := dto.LoginResponse{
		Email: userRequest.Email,
		Token: token,
	}

	response := helper.ResponseData(http.StatusOK, "Login Successfully!", responseUser)
	return c.JSON(http.StatusOK, response)
}
