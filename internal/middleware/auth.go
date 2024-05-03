package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secretKey := config.GetConfig().Server.JWTSecret

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is not provided")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token format. Use Bearer token")
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signature")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if claims, ok := token.Claims.(*helper.JwtCustomClaims); ok && next != nil {
			if claims.Role != "user" && claims.Role != "admin" {
				return echo.ErrUnauthorized
			}
			c.Set("user", claims)
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
