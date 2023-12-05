package orchestration

import (
	"context"
	"fmt"
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
	fmt.Println("two **")
	return g.GptService.SendMessage(ctx, message)
}
