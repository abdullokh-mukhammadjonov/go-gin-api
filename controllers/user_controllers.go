package controllers

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/database"
	"github.com/abdullokh-mukhammadjonov/gin-api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func CreateUser(ctx *gin.Context) {
	var input CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(401, gin.H{"status": "fail", "error": err.Error()})
	}

	user := models.Person{Name: input.Name, Surname: input.Surname}

	database.DB.Create(&user)
	ctx.JSON(201, gin.H{"status": "success", "user": user})
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.Person

	database.DB.Find(&users)

	ctx.JSON(200, gin.H{"satus": "success", "users": users})
}
