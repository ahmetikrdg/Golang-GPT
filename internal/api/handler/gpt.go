package handler

import (
	"github.com/gin-gonic/gin"

	"trendyol-gpt/internal/api/dto/request"
	"trendyol-gpt/internal/api/orchestration"
)

type IChatHandler interface {
	SendMessage(c *gin.Context)
}

type ChatHandler struct {
	Orchestration orchestration.GptOrchestration
}

func (ch ChatHandler) SendMessage(c *gin.Context) {
	var msg request.MessageRequest

	if err := c.ShouldBindJSON(&msg); err != nil {
		return
	}

	result, err := ch.Orchestration.SendMessage(c, msg.Message)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": result})
}
