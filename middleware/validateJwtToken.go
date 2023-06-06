package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJwtToken(ctx *gin.Context) {
	// Get the value of the "Authorization" header
	authHeader := ctx.GetHeader("Authorization")

	// Check if the authorization header exists and starts with "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		return
	}

	// Extract the token from the authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// decode and validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verifying method used for token signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		return hmacSecret, nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expiry
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// attach to request
		ctx.Set("userEmail", claims["email"])
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
