package controllers

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/database"
	"github.com/abdullokh-mukhammadjonov/gin-api/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title     string `json:"title" binding:"required"`
	AuthorID  uint64 `json:"author" binding:"required"`
	WrittenIn string `json:"written_in" binding:"required"`
}

func GetBooks(ctx *gin.Context) {
	var books []models.Book
	database.DB.Preload("Author").Find(&books)
	ctx.JSON(200, gin.H{"message": "success", "data": books})
}

func CreateBook(ctx *gin.Context) {
	// validating input
	var inputData CreateBookInput

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(401, gin.H{"status": "fail", "error": err.Error()})
	}

	// create book
	book := models.Book{Title: inputData.Title, AuthorID: inputData.AuthorID, WrittenIn: inputData.WrittenIn}

	database.DB.Create(&book).Preload("Author")

	ctx.JSON(201, gin.H{"status": "success", "book": book})
}
