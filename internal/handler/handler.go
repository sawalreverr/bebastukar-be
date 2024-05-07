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
	GetAllDiscussionFromProfile(c echo.Context) error
	NewDiscussionHandler(c echo.Context) error
	EditDiscussionhandler(c echo.Context) error
	DeleteDiscussionhandler(c echo.Context) error
	FindAllDiscussionUserHandler(c echo.Context) error
	FindDiscussionByID(c echo.Context) error
}
