package main

import (
	"net/http"

	"example.com/restapi/eventbooking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // this will return engine means a http server
	server.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello WOrld!!!"})
	}) // GET, POST, PUT, DELETE, PATCH

	server.GET("/events", getEvents)
	server.POST("/events", saveEvent)
	server.Run(":8081")
}

func getEvents(ctx *gin.Context) {
	events := models.GetEvents()
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func saveEvent(ctx *gin.Context) {
	var event models.Event
	ctx.ShouldBindJSON(&event)
	ctx.JSON(http.StatusAccepted, gin.H{"message": "created"})
}
