package controllers

import (
	"HackOn/models"
	"HackOn/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterChatRoutes(r *gin.Engine) {
	r.POST("/chat", PostChatMessage)
}

func SortedChatRoutes(r *gin.Engine) {
	r.GET("/chat/sorted", GetChatMessagesSorted)
}

func PostChatMessage(c *gin.Context) {
	var msg models.ChatMessage
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	msg.Timestamp = time.Now().Unix()

	if err := repositories.CreateChatMessage(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func GetChatMessagesSorted(c *gin.Context) {
	messages, err := repositories.GetChatMessagesSorted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
