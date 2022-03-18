package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// uuid "github.com/satori/go.uuid"
	"github.com/google/uuid"
	"github.com/vitorscassiano/todoapp/models"
	"github.com/vitorscassiano/todoapp/wire"
)

func FindListByUserID(ctx *gin.Context) {
	var lists []models.List
	userID := ctx.Param("id")

	models.DB.Where("UserRef = ?", userID).Find(&lists)

	ctx.JSON(http.StatusOK, gin.H{"lists": lists})
}

func CreateListByUserID(ctx *gin.Context) {
	var input wire.CreateList

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID := ctx.Param("id")
	parsedUUID, err := uuid.Parse(userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	lists := models.List{
		Name:        input.Name,
		Description: input.Description,
		UserRef:     parsedUUID,
	}

	models.DB.Create(&lists)

	ctx.JSON(http.StatusOK, gin.H{
		"user":  userID,
		"lists": lists,
	})
}
