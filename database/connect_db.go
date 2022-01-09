package database

import (
	"github.com/abdullokh-mukhammadjonov/gin-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=abdulloh dbname=book_store password=333333clean")

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Book{}, &models.Person{}, &models.Task{})

	DB = database
}
