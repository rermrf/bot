package main

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"os"
)

func main() {
	token := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(token)
	ctx := context.Background()

	const approxTokensPerWord = 4
	const wordCount = 100
	requiredTokens := approxTokensPerWord * wordCount

	var historyMessages []openai.ChatCompletionMessage
	var content string
	for {
		fmt.Print("：")
		_, err := fmt.Scanln(&content)
		if err != nil {
			return
		}
		message := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		}

		historyMessages = append(historyMessages, message)

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: requiredTokens, // Updated to accommodate a 1000-word story
			Messages:  historyMessages,
			Stream:    true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()

		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println()
				break
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				return
			}

			for _, message := range response.Choices {
				fmt.Printf(message.Delta.Content)
			}
		}
	}
}
