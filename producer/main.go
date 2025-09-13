package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

type RequestBody struct {
	Message string `json:"message" binding:"required"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/event", func(c *gin.Context) {
		var message RequestBody

		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received message: %s\n", message.Message)

		c.String(http.StatusOK, "Message received")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
