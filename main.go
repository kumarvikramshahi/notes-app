package main

import (
	"log"
	"notesApp/configs"
	"notesApp/routes"

	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load env variables
	err:=godotenv.Load()
	if err!=nil {
		log.Fatal("not able to load env variables")
	}
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
