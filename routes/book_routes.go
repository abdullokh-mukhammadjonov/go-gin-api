package routes

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func bookRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/books")
	{
		booksRouter.GET("/", controllers.GetBooks)

		booksRouter.POST("/", controllers.CreateBook)
	}
}
