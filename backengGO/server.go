package main

import (
	"net/http"

	"github.com/Baral-Chief-of-Compliance/API_GAME_2.0/backendGO/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})

	server.GET("/users", controller.GetAllUsers)
	server.POST("/users", controller.AddUser)
	server.DELETE("/users/:id", controller.DeleteUser)
	server.PUT("/users/:id", controller.UpdateUser)

	server.Run(":8080")

}
