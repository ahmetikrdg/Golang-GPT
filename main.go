package main

import (
	"github.com/gin-gonic/gin"
	"trendyol-gpt/internal/api"
	gpt2 "trendyol-gpt/internal/api/gpt"
	handler2 "trendyol-gpt/internal/api/handler"
	"trendyol-gpt/internal/api/orchestration"
	chatGpt "trendyol-gpt/pkg/chat-gpt"
)

func main() {
	gpt := chatGpt.ChatGpt{}
	service := gpt2.NewGptService(gpt)
	orc := orchestration.NewGptOrchestration(*service)
	handler := handler2.ChatHandler{Orchestration: *orc}

	app := gin.Default()
	api.NewRouteSetup(app, handler)

	err := app.Run(":8888")
	if err != nil {
		return
	}
}
