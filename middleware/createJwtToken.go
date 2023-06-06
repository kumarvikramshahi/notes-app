package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(email string) (string, error) {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"expiry": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
