package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
)

type discussionHandler struct {
	disccusionUsecase usecase.DiscussionUsecase
}

func NewDiscussionHandler(uc usecase.DiscussionUsecase) DiscussionHandler {
	return &discussionHandler{disccusionUsecase: uc}
}

func (h *discussionHandler) GetAllDiscussionFromProfile(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	discussions, err := h.disccusionUsecase.GetAllDiscussionFromUser(userID)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	if len(*discussions) == 0 {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion empty")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusFound, "found!", discussions))
}

func (h *discussionHandler) NewDiscussionHandler(c echo.Context) error {
	var discussionData dto.DiscussionInput

	claims := c.Get("user").(*helper.JwtCustomClaims)
	userID := claims.UserID

	if err := c.Bind(&discussionData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error")
	}

	if err := c.Validate(&discussionData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "content must not be empty and must have at least 1 letter")
	}

	form, _ := c.MultipartForm()
	imageFiles := form.File["images"]

	validImages, err := helper.ImagesValidation(imageFiles)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	var urls []string
	for _, file := range validImages {
		resultURL, _ := helper.UploadToCloudinary(file, "bebastukar/discussion/")
		urls = append(urls, resultURL)
	}

	dataCredential := dto.DiscussionCredential{
		AuthorID: userID,
		Content:  discussionData.Content,
		Images:   urls,
	}

	response, err := h.disccusionUsecase.CreateDiscussion(dataCredential)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusCreated, helper.ResponseData(http.StatusCreated, "discussion created!", response))
}

func (h *discussionHandler) EditDiscussionhandler(c echo.Context) error {
	var discussionData dto.DiscussionInput

	discussionID := c.Param("id")

	if err := c.Bind(&discussionData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&discussionData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "content must not be empty and must have at least 1 letter")
	}

	discussionFound, err := h.disccusionUsecase.GetDiscussionFromID(discussionID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion not found!")
	}

	dataCredential := dto.DiscussionCredential{
		AuthorID: discussionFound.AuthorID,
		Content:  discussionData.Content,
	}

	if err := h.disccusionUsecase.EditDiscussion(discussionID, dataCredential); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "discussion updated!", nil))
}

func (h *discussionHandler) DeleteDiscussionhandler(c echo.Context) error {
	discussionID := c.Param("id")

	discussionFound, err := h.disccusionUsecase.GetDiscussionFromID(discussionID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion not found!")
	}

	if err := h.disccusionUsecase.DeleteDiscussion(discussionFound.ID); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error!")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "discussion deleted!", nil))
}

func (h *discussionHandler) FindDiscussionByID(c echo.Context) error {
	discussionID := c.Param("id")

	discussionFound, err := h.disccusionUsecase.GetDiscussionFromID(discussionID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion not found!")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "found!", discussionFound))
}

func (h *discussionHandler) FindAllDiscussionUserHandler(c echo.Context) error {
	userID := c.Param("userid")

	discussionFound, err := h.disccusionUsecase.GetAllDiscussionFromUser(userID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion not found!")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "found!", discussionFound))
}

func (h *discussionHandler) FindAllDiscussion(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	sortBy := c.QueryParam("sort_by")
	sortType := c.QueryParam("sort_type")

	if sortBy == "" {
		sortBy = "created_at"
		sortType = "desc"
	}

	if sortType == "" {
		sortType = "asc"
	}

	discussionResponse, err := h.disccusionUsecase.GetAllDiscussion(page, limit, sortBy, sortType)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "OK", discussionResponse))
}
