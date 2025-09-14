package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twmb/franz-go/pkg/kgo"
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

		produceMessage(message.Message)

		c.String(http.StatusOK, "Message received")
	})

	return r
}

func produceMessage(message string) {
	opts := []kgo.Opt{
		kgo.SeedBrokers("localhost:9092"),
		kgo.DefaultProduceTopic("screams"),
		kgo.ClientID("my-client-id"),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		// TODO: handle/log this error
		return
	}
	defer client.Close()

	record := &kgo.Record{
		Value: []byte(message),
		Topic: "screams",
	}

	fmt.Println("Producing message to Kafka...")

	if err := client.ProduceSync(context.Background(), record).FirstErr(); err != nil {
		fmt.Printf("Error producing message: %v\n", err)
		return
	}
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
