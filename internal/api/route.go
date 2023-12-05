package api

import (
	"github.com/gin-gonic/gin"
	"trendyol-gpt/internal/api/handler"
)

func NewRouteSetup(app *gin.Engine, handler handler.IChatHandler) {

	app.POST("/message", handler.SendMessage)
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
