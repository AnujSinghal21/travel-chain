package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your_secret_key") // Replace with strong key

// Claims structure for JWT
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates JWT token for given user
func GenerateToken(email, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
