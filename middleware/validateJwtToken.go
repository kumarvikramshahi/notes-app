package middleware

import (
	"fmt"
	"net/http"
	"notesApp/configs"
	"notesApp/schemas"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJwtToken(ctx *gin.Context) {
	tokenString := ctx.Param("sessionId")

	// decode and validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		return hmacSampleSecret, nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expiry
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// find user with in token
		user := &schemas.User{}
		result := configs.DefaultDB.First(&user, claims["email"])
		if result != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach to request
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
