package chat_gpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type IGpt interface {
	SendMessage(ctx context.Context, message string) (string, error)
}

type ChatGpt struct {
}

func NewChatGptServer() *ChatGpt {
	return &ChatGpt{}
}

func (c ChatGpt) SendMessage(ctx context.Context, message string) (string, error) {
	client := openai.NewClient("key************")

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	fmt.Println("result : ", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil
}
