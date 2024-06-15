package main

import (
	"task_manager/back/database"
	"task_manager/back/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	routes.SetupRoutes(router)
	router.Run(":8080")

}
