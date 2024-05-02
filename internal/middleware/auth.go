package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
)

func NewJWTMiddleware() echo.MiddlewareFunc {
	secretKey := config.GetConfig().Server.JWTSecret
	configJWT := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	}

	return echojwt.WithConfig(configJWT)
}

func UserValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(*jwt.Token).Claims.(*helper.JwtCustomClaims)
		fmt.Println(claims)
		if claims.Role != "user" && claims.Role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func AdminValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(*jwt.Token).Claims.(*helper.JwtCustomClaims)
		fmt.Println(claims)
		if claims.Role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}

}
