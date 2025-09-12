package main

import (
	"net/http"
	"time"

	"com.chatinho/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	chats = make(map[uuid.UUID]*models.Chat)
)

func main() {

	router := gin.Default()

	/// 1. Create send message [POST]

	router.POST("/send-message", func(context *gin.Context) {

		body := models.Message{}

		err := context.ShouldBindJSON(&body)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		messageId := uuid.New()

		if body.ID == uuid.Nil {

			chatID := uuid.New()
			CreatedAt := time.Now()

			chats[chatID] = &models.Chat{
				ID:        chatID,
				CreatedAt: CreatedAt,
				Messages: []models.Message{{
					ID:      messageId,
					Message: body.Message,
					Time:    time.Now(),
					User:    body.User,
				}},
			}
			context.JSON(http.StatusAccepted, gin.H{
				"chat_id":    chatID,
				"message_id": messageId,
			})
			context.Abort()
			return

		} else if chat, exists := chats[body.ID]; exists {
			chat.Messages = append(chat.Messages, models.Message{
				ID:      messageId,
				Message: body.Message,
				Time:    time.Now(),
				User:    body.User,
			})

			context.JSON(http.StatusAccepted, &messageId)
			context.Abort()
			return
		}
		context.JSON(http.StatusNotFound, gin.H{"Error": "Chat not found"})
	})

	/// 2. Create list chats [GET]

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, chats)
	})

	/// 3. Delete message [DELETE]

	/// 4. Update message [PATH]

	router.Run()
}
