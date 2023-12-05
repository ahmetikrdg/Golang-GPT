package gpt

import (
	"context"

	chatGpt "trendyol-gpt/pkg/chat-gpt"
)

type IGptService interface {
	SendMessage(ctx context.Context, message string) (string, error)
}

type GptService struct {
	chatGpt chatGpt.ChatGpt
}

func NewGptService(cg chatGpt.ChatGpt) *GptService {
	return &GptService{chatGpt: cg}
}

func (g GptService) SendMessage(ctx context.Context, message string) (string, error) {
	return g.chatGpt.SendMessage(ctx, message)
}
