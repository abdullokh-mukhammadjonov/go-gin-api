package routes

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func taskRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/tasks")
	{
		booksRouter.GET("/", controllers.GetAllTasks)

		booksRouter.POST("/", controllers.CreateTask)
	}
}
