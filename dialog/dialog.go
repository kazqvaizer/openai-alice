package dialog

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type DialogConfig struct {
	ApiKey string
}

func AskAlice(question string, config DialogConfig) (string, error) {

	client := openai.NewClient(config.ApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0125,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
