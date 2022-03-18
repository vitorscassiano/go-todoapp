package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/todoapp/controllers"
	"github.com/vitorscassiano/todoapp/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	// Version Controller
	router.GET("api/version", func(ctx *gin.Context) {
		response := map[string]string{
			"name":    "ToDo App",
			"version": "0.0.1",
		}

		ctx.JSON(http.StatusOK, gin.H{"data": response})
	})
	// Controllers

	router.GET("api/users", controllers.FindUsers)
	router.POST("api/users", controllers.CreateUser)

	router.GET("api/users/:id/lists", controllers.FindListByUserID)
	router.POST("api/users/:id/lists", controllers.CreateListByUserID)

	router.Run()
}
