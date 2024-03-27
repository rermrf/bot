package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

func main() {
	key := os.Getenv("OPENAI_API_KEY")
	fmt.Println(key)
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo16K,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Chat completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
