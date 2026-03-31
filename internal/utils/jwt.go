package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Cliams struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userUD string, secret string, duration time.Duration) (string, error) {
	claims := Cliams{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}
