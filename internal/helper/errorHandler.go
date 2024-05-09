package helper

import "github.com/labstack/echo/v4"

func ErrorHandler(c echo.Context, statusCode int, errorMessage string) error {
	response := ResponseData(statusCode, errorMessage, nil)
	return c.JSON(statusCode, response)
}
