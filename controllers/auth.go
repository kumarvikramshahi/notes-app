package controllers

import (
	"net/http"
	"net/mail"
	"notesApp/configs"
	"notesApp/middleware"
	"notesApp/schemas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func isValidMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func LogIn(ctx *gin.Context) {
	requestData := &schemas.User{}
	// bind incoming JSON to struct
	if err := ctx.ShouldBindJSON(requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !isValidMailAddress(requestData.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email adress type",
		})
		return
	}

	user := &schemas.User{}
	result := configs.DefaultDB.Where("email = ?", requestData.Email).First(user)
	// if user not found
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "please register, user not found",
		})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	// if email found but password incorrect
	if user.Password != requestData.Password {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "either email or password wrong",
		})
		return
	}
	// if user found & password correct => return session_id
	jwtToken, err := middleware.CreateJwtToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "can't create token",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sid": jwtToken,
	})
}

func SignUp(ctx *gin.Context) {
	user := &schemas.User{}
	// bind incoming JSON to struct
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !isValidMailAddress(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email adress type",
		})
		return
	}

	// check if user exist
	result := configs.DefaultDB.First(user)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// user doesn't exist => create user
		if result := configs.DefaultDB.Create(user); result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	// user already exist
	ctx.JSON(http.StatusAlreadyReported, gin.H{
		"message": "user with mentioned email id already exist",
	})

}
