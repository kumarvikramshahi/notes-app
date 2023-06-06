package routes

import (
	"notesApp/controllers"
	"notesApp/middleware"

	"github.com/gin-gonic/gin"
)

func Notes(router *gin.Engine) {
	router.GET("/notes/:sessionId", middleware.ValidateJwtToken, controllers.FetchNotes)
	router.POST("/notes", middleware.ValidateJwtToken, controllers.CreateNotes)
	router.DELETE("/notes/:nid/:sid", middleware.ValidateJwtToken, controllers.DeleteNotes)
}
