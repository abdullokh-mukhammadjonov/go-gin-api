package routes

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/users")
	{
		booksRouter.GET("/", controllers.GetAllUsers)

		booksRouter.POST("/", controllers.CreateUser)
	}
}
