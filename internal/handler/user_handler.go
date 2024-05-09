package handler

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uc}
}

func (h *userHandler) ProfileGet(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	userFound, err := h.userUsecase.FindUserByID(userID)
	if err != nil {
		internalErr := helper.ResponseData(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	userResponse := dto.ProfileUser{
		ID:          userFound.ID,
		Name:        userFound.Name,
		Email:       userFound.Email,
		PhoneNumber: userFound.PhoneNumber,
		Role:        userFound.Role,
		ImageURL:    userFound.ImageURL,
		Bio:         userFound.Bio,
	}

	response := helper.ResponseData(http.StatusOK, "Your profile information!", userResponse)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) ProfileUpdate(c echo.Context) error {
	var userRequest dto.UpdateUser

	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	if err := c.Bind(&userRequest); err != nil {
		bindErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	if err := c.Validate(&userRequest); err != nil {
		validateErr := helper.ResponseData(http.StatusBadRequest, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, validateErr)
	}

	if err := h.userUsecase.UpdateUser(userID, userRequest); err != nil {
		internalErr := helper.ResponseData(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	response := helper.ResponseData(http.StatusOK, "Update Successfully!", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) FindUser(c echo.Context) error {
	userID := c.Param("id")
	id, err := uuid.Parse(userID)

	if err != nil {
		parseErr := helper.ResponseData(http.StatusBadRequest, "ID invalid", nil)
		return c.JSON(http.StatusBadRequest, parseErr)
	}

	userFound, err := h.userUsecase.FindUserByID(id.String())

	if err != nil {
		internalErr := helper.ResponseData(http.StatusInternalServerError, err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, internalErr)
	}

	userResponse := dto.ProfileUser{
		ID:          userFound.ID,
		Name:        userFound.Name,
		Email:       userFound.Email,
		PhoneNumber: userFound.PhoneNumber,
		Role:        userFound.Role,
		ImageURL:    userFound.ImageURL,
		Bio:         userFound.Bio,
	}

	response := helper.ResponseData(http.StatusFound, "User Found!", userResponse)
	return c.JSON(http.StatusFound, response)
}

func (h *userHandler) UploadAvatar(c echo.Context) error {
	file, err := c.FormFile("image")

	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	if err != nil {
		missingErr := helper.ResponseData(http.StatusBadRequest, "please upload your image!", nil)
		return c.JSON(http.StatusBadRequest, missingErr)
	}

	if file.Size > 2*1024*1024 {
		tooLarge := helper.ResponseData(http.StatusBadRequest, "upload image size must less than 2MB!", nil)
		return c.JSON(http.StatusBadRequest, tooLarge)
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		typeErr := helper.ResponseData(http.StatusBadRequest, "only image allowed!", nil)
		return c.JSON(http.StatusBadRequest, typeErr)
	}

	src, _ := file.Open()
	defer src.Close()

	resp, err := helper.UploadToCloudinary(src, "bebastukar/avatar/")
	if err != nil {
		cloudUploadErr := helper.ResponseData(http.StatusInternalServerError, "upload error!", nil)
		return c.JSON(http.StatusInternalServerError, cloudUploadErr)
	}

	err = h.userUsecase.UpdateUserAvatar(userID, resp)
	if err != nil {
		updateErr := helper.ResponseData(http.StatusInternalServerError, "update database error!", nil)
		return c.JSON(http.StatusInternalServerError, updateErr)
	}

	response := helper.ResponseData(http.StatusOK, "Image Uploaded!", nil)
	return c.JSON(http.StatusOK, response)
}
