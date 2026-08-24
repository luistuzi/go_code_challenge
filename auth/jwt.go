package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(user string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
