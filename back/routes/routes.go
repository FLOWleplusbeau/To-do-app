package routes

import (
	"task_manager/back/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
	}
}
