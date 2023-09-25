package controller

import (
	"log"
	"net/http"

	"github.com/Baral-Chief-of-Compliance/API_GAME_2.0/backendGO/database"
	"github.com/Baral-Chief-of-Compliance/API_GAME_2.0/backendGO/model"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {

	users, err := model.AllUsers(database.DB)

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}

func AddUser(context *gin.Context) {
	var user model.UserEnter
	context.BindJSON(&user)

	result, err := model.AddUser(database.DB, user.Name, user.PassHash)

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	result, err := model.DeleteUser(database.DB, id)

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}

func UpdateUser(context *gin.Context) {
	id := context.Param("id")

	var user model.User

	context.BindJSON(&user)

	result, err := model.UpdateUser(database.DB, id, user)

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
