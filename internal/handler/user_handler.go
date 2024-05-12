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
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
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

	response := helper.ResponseData(http.StatusOK, "ok", userResponse)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) ProfileUpdate(c echo.Context) error {
	var userRequest dto.UpdateUser

	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	if err := c.Bind(&userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "name or phone number invalid!")
	}

	if err := h.userUsecase.UpdateUser(userID, userRequest); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
	}

	response := helper.ResponseData(http.StatusOK, "update successfully!", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) FindUser(c echo.Context) error {
	userID := c.Param("id")
	id, err := uuid.Parse(userID)

	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "param id invalid!")
	}

	userFound, err := h.userUsecase.FindUserByID(id.String())

	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "user not found!")
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

	response := helper.ResponseData(http.StatusFound, "ok", userResponse)
	return c.JSON(http.StatusFound, response)
}

func (h *userHandler) UploadAvatar(c echo.Context) error {
	file, err := c.FormFile("image")

	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "please upload your image!")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB!")
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed!")
	}

	src, _ := file.Open()
	defer src.Close()

	resp, err := helper.UploadToCloudinary(src, "bebastukar/avatar/")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "upload failed, cloudinary server error!")
	}

	err = h.userUsecase.UpdateUserAvatar(userID, resp)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "update database error!")
	}

	response := helper.ResponseData(http.StatusOK, "image uploaded!", nil)
	return c.JSON(http.StatusOK, response)
}
