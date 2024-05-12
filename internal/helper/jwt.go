package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sawalreverr/bebastukar-be/config"
)

type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userID string, role string) (string, error) {
	claims := &JwtCustomClaims{
		userID, role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.GetConfig().Server.JWTSecret
	return token.SignedString([]byte(secret))
}
