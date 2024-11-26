package utils

import (
	"time"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("mysecretkey")

type JWTClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *user.User) (string, error) {
	claims := &JWTClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)

	return signed, err
}
