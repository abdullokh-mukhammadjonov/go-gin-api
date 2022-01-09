package main

import (
	"fmt"

	"github.com/abdullokh-mukhammadjonov/gin-api/database"
	"github.com/abdullokh-mukhammadjonov/gin-api/routes"
	"github.com/gin-gonic/gin"
)

// taking a struct
type Rect struct {
	len, wid int
}

func (re Rect) Area() int {
	return re.len * re.wid
}

func (re Rect) Area_by_value() int {
	return re.len * re.wid
}

func (re *Rect) Area_by_reference() int {
	return re.len * re.wid
}

func startServer() {
	app := gin.New()
	database.ConnectDatabase()

	app.GET("/api", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "success", "message": "api is running"}) })
	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	defer database.DB.Close()

	app.Run(":5500")
}

// main function
func main() {
	startServer()
	r := Rect{10, 12}
	fmt.Println("Length and Width is:", r)
	fmt.Println("Area of Rectangle is:", r.Area_by_value())
	fmt.Println("Area of Rectangle is:", r.Area_by_reference())
	fmt.Println("Area of Rectangle is:", (&r).Area_by_value())
	fmt.Println("Area of Rectangle is:", (&r).Area_by_reference())
}
