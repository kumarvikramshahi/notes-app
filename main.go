package main

import (
	"notesApp/configs"
	"notesApp/routes"

	gin "github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectDB()
}

func main() {
	appRouter := gin.Default()
	appRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	// authentication related routes
	routes.Auth(appRouter)

	// notes routes
	routes.Notes(appRouter)

	// run project
	appRouter.Run()
}
