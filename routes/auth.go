package routes

import (
	"notesApp/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.Engine) {
	router.POST("/login", controllers.LogIn)
	router.POST("/signup", controllers.SignUp)
}
