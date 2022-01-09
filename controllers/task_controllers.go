package controllers

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/database"
	"github.com/abdullokh-mukhammadjonov/gin-api/models"
	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	Name      string `json:"name" binding:"required"`
	Status    string `json:"status"`
	Priority  uint64 `json:"priority"`
	CreatorID uint64 `json:"created_by" binding:"required"`
	DueDate   string `json:"due_date"`
}

func CreateTask(ctx *gin.Context) {
	var input createTaskRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"status": "fail", "error": err.Error()})
	}

	task := models.Task{
		Name:      input.Name,
		Status:    input.Status,
		Priority:  input.Priority,
		CreatorID: input.CreatorID,
		DueDate:   input.DueDate,
	}

	database.DB.Create(&task)

	ctx.JSON(201, gin.H{"status": "success", "task": task})
}

func GetAllTasks(ctx *gin.Context) {
	var tasks []models.Task

	database.DB.Preload("CreatedBy").Find(&tasks)

	ctx.JSON(200, gin.H{"status": "success", "tasks": tasks})
}
