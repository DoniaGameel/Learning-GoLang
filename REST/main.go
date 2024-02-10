package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
	"example/rest/models"
	"fmt"
)

func main() {
	// configure an HTTP server
    router := gin.Default()
	// setting a handler for incoming GET request
    router.GET("/events", getEvents)
	// setting a handler for incoming POST request
	router.POST("/events", createEvents)
	//listen on port 8080
    router.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	// send back a response in JSON format
	//context.JSON(http.StatusOK, gin.H{"message":"Hello!"})
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context){
	var event models.Event
	// the data from request body will be stored in that variable
	err := context.ShouldBindJSON(&event)
	if err!= nil{
		fmt.Print(err)
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data."})
	}

	// just a dummy value till we use a database
	else{
		event.ID = 1
		event.UserID = 1

		event.Save()
		context.JSON(http.StatusCreated,gin.H{"message":"Event Created!", "event": event})
	}
	}