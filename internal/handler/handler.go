package handler

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	RegisterHandler(c echo.Context) error
	LoginHandler(c echo.Context) error
}

type UserHandler interface {
	ProfileGet(c echo.Context) error
	ProfileUpdate(c echo.Context) error
	UploadAvatar(c echo.Context) error
	FindUser(c echo.Context) error
}

type DiscussionHandler interface {
	// Discussion handler
	GetAllDiscussionFromProfile(c echo.Context) error
	NewDiscussionHandler(c echo.Context) error
	EditDiscussionhandler(c echo.Context) error
	DeleteDiscussionhandler(c echo.Context) error

	// Discussion finder handler
	FindAllDiscussionUserHandler(c echo.Context) error
	FindDiscussionByID(c echo.Context) error
	FindAllDiscussion(c echo.Context) error

	// Discussion comment handler
	AddDiscussionCommentHandler(c echo.Context) error
	EditDiscussionCommentHandler(c echo.Context) error
	DeleteDiscussionCommentHandler(c echo.Context) error

	// Discussion reply comment handler
	AddDiscussionReplyCommentHandler(c echo.Context) error
	EditDiscussionReplyCommentHandler(c echo.Context) error
	DeleteDiscussionReplyCommentHandler(c echo.Context) error
}
