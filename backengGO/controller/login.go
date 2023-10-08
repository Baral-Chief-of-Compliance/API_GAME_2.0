package controller

import (
	"fmt"
	"net/http"

	"github.com/Baral-Chief-of-Compliance/API_GAME_2.0/backendGO/database"
	"github.com/Baral-Chief-of-Compliance/API_GAME_2.0/backendGO/model"
	"github.com/gin-gonic/gin"
)

func Enter(context *gin.Context) {
	var user model.UserEnter
	var statusCode int = http.StatusOK
	context.BindJSON(&user)
	passHash, err := model.GetInfUserLogin(database.DB, user.Login)

	if err != nil {
		statusCode = http.StatusUnauthorized
	}
	fmt.Print(passHash)

	context.JSON(statusCode, gin.H{
		"message": "da",
	})
}
