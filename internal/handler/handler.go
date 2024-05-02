package handler

import "github.com/labstack/echo/v4"

type UserHandler interface {
	RegisterHandler(c echo.Context) error
	LoginHandler(c echo.Context) error
}
