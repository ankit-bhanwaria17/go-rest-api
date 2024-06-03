package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvents)

	server.Run(":8080")
}

func GetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func CreateEvents(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err})
		return
	}
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
