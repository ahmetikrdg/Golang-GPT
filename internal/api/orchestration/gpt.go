package orchestration

import (
	"context"
	"trendyol-gpt/internal/api/gpt"
)

type IGptOrchestration interface {
	SendMessage(ctx context.Context, message string) (string, error)
}

type GptOrchestration struct {
	GptService gpt.GptService
}

func NewGptOrchestration(gs gpt.GptService) *GptOrchestration {
	return &GptOrchestration{
		GptService: gs,
	}
}

func (g GptOrchestration) SendMessage(ctx context.Context, message string) (string, error) {
	return g.GptService.SendMessage(ctx, message)
}
