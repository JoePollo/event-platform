package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"message": "hello world",
		},
	)
}

func connectToDb(ctx *gin.Context) {
	dsn := "host=localhost user=eventuser password=eventpass dbname=event_service port=5432 sslmode=disabled timezone=US/Chicago"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Failed to connect to database due to error:\n %v.", err)})
	}
	ctx.JSON(
		http.StatusOK, gin.H{"message": "Successfully connected to the database!"},
	)
}

func main() {
	router := gin.Default()
	router.GET("/helloWorld", helloWorld)
	router.POST("/connectToDb", connectToDb)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
