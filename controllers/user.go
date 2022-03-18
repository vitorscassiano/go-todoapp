package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/todoapp/models"
	"github.com/vitorscassiano/todoapp/wire"
)

func FindUsers(ctx *gin.Context) {
	var users []models.User
	var lists []models.List

	models.DB.Find(&users)
	models.DB.Find(&lists)

	for _, user := range users {
		for _, list := range lists {
			if list.UserRef == user.UserID {
				user.Lists = append(user.Lists, list)
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(ctx *gin.Context) {
	var input wire.CreateUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{
		Name: input.Name,
	}

	models.DB.Create(&user)

	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}
