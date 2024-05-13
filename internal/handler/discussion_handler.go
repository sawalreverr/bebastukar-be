package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
	"github.com/sawalreverr/bebastukar-be/pkg"
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

	return c.JSON(http.StatusFound, helper.ResponseData(http.StatusFound, "ok", discussions))
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

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "ok", discussionFound))
}

func (h *discussionHandler) FindAllDiscussionUserHandler(c echo.Context) error {
	userID := c.Param("userID")

	discussionFound, err := h.disccusionUsecase.GetAllDiscussionFromUser(userID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusNotFound, "discussion not found!")
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "ok", discussionFound))
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

func (h *discussionHandler) AddDiscussionCommentHandler(c echo.Context) error {
	var commentData dto.DiscussionCommentInput

	discussionID := c.Param("id")

	if err := c.Bind(&commentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&commentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "comment must not be empty and must have at least 1 letter")
	}

	claims := c.Get("user").(*helper.JwtCustomClaims)

	commentCred := dto.DiscussionCommentCredential{
		AuthorID:     claims.UserID,
		DiscussionID: discussionID,
		Comment:      commentData.Comment,
	}

	response, err := h.disccusionUsecase.CreateCommentDiscussion(commentCred)
	if err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusCreated, helper.ResponseData(http.StatusCreated, "comment added!", response))
}

func (h *discussionHandler) EditDiscussionCommentHandler(c echo.Context) error {
	var commentData dto.DiscussionCommentInput

	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")

	if err := c.Bind(&commentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&commentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "comment must not be empty and must have at least 1 letter")
	}

	claims := c.Get("user").(*helper.JwtCustomClaims)

	commentCred := dto.DiscussionCommentCredential{
		AuthorID:     claims.UserID,
		DiscussionID: discussionID,
		Comment:      commentData.Comment,
	}

	if err := h.disccusionUsecase.EditCommentDiscussion(discussionCommentID, commentCred); err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "comment updated!", nil))
}

func (h *discussionHandler) DeleteDiscussionCommentHandler(c echo.Context) error {
	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := h.disccusionUsecase.DeleteCommentDiscussion(discussionID, discussionCommentID, claims.UserID); err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "comment deleted!", nil))
}

func (h *discussionHandler) FindAllDiscussionCommentHandler(c echo.Context) error {
	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")

	response, err := h.disccusionUsecase.GetAllCommentFromDiscussionPublic(discussionID, discussionCommentID)
	if err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "ok", response))
}

func (h *discussionHandler) AddDiscussionReplyCommentHandler(c echo.Context) error {
	var replyCommentData dto.DiscussionCommentInput

	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := c.Bind(&replyCommentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&replyCommentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "comment must not be empty and must have at least 1 letter")
	}

	replyCred := dto.DiscussionReplyCommentCredential{
		AuthorID:            claims.UserID,
		DiscussionID:        discussionID,
		DiscussionCommentID: discussionCommentID,
		ReplyComment:        replyCommentData.Comment,
	}

	response, err := h.disccusionUsecase.CreateReplyCommentDiscussion(replyCred)
	if err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusCreated, helper.ResponseData(http.StatusCreated, "reply comment created!", response))
}

func (h *discussionHandler) EditDiscussionReplyCommentHandler(c echo.Context) error {
	var replyCommentData dto.DiscussionCommentInput

	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")
	discussionReplyCommentID := c.Param("replyCommentID")
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := c.Bind(&replyCommentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error!")
	}

	if err := c.Validate(&replyCommentData); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "comment must not be empty and must have at least 1 letter")
	}

	replyCred := dto.DiscussionReplyCommentCredential{
		AuthorID:            claims.UserID,
		DiscussionID:        discussionID,
		DiscussionCommentID: discussionCommentID,
		ReplyComment:        replyCommentData.Comment,
	}

	if err := h.disccusionUsecase.EditReplyCommentDiscussion(discussionReplyCommentID, replyCred); err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "reply comment updated!", nil))
}

func (h *discussionHandler) DeleteDiscussionReplyCommentHandler(c echo.Context) error {
	discussionID := c.Param("id")
	discussionCommentID := c.Param("commentID")
	discussionReplyCommentID := c.Param("replyCommentID")
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := h.disccusionUsecase.DeleteReplyCommentDiscussion(discussionID, discussionCommentID, discussionReplyCommentID, claims.UserID); err != nil {
		return pkg.DiscussionErrorHelper(c, err)
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "reply comment deleted!", nil))
}
