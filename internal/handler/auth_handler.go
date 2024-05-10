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
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "validation error!")
	}

	newUser, err := h.userUsecase.RegisterUser(userRequest)

	if err != nil {
		if errors.Is(err, pkg.ErrDataAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusConflict, "email already registered!")
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
	}

	responseUser := dto.RegisterResponse{
		ID:          newUser.ID,
		Name:        newUser.Name,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
	}

	response := helper.ResponseData(http.StatusCreated, "register successfully!", responseUser)
	return c.JSON(http.StatusCreated, response)
}

func (h *authHandler) LoginHandler(c echo.Context) error {
	var userRequest dto.Login

	if err := c.Bind(&userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "validation error!")
	}

	token, err := h.userUsecase.LoginUser(userRequest.Email, userRequest.Password)

	if err != nil {
		if errors.Is(err, pkg.ErrRecordNotFound) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "email or password invalid!")
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
	}

	responseUser := dto.LoginResponse{
		Email: userRequest.Email,
		Token: token,
	}

	response := helper.ResponseData(http.StatusOK, "login successfully!", responseUser)
	return c.JSON(http.StatusOK, response)
}
