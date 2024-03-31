package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func Chat(usermessage string, client chan []byte) error {
	if usermessage == "" {
		return errors.New("empty message")
	}
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		return errors.New("OPENAI_API_KEY environment variable not set")
	}
	c := openai.NewClient(token)
	ctx := context.Background()

	const approxTokensPerWord = 4
	const wordCount = 100
	requiredTokens := approxTokensPerWord * wordCount

	var historyMessages []openai.ChatCompletionMessage

	message := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: usermessage,
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
		return fmt.Errorf("ChatCompletionStream error: %v\n", err)
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		//if errors.Is(err, io.EOF) {
		//	fmt.Println()
		//	return nil
		//}

		if err != nil {
			return fmt.Errorf("\nStream error: %v\n", err)
		}

		for _, message := range response.Choices {
			fmt.Printf(message.Delta.Content)
			client <- []byte(message.Delta.Content)
		}
	}

}

//func Chat1(usermessage string) error {
//	if usermessage == "" {
//		return errors.New("empty message")
//	}
//	token := os.Getenv("OPENAI_API_KEY")
//	if token == "" {
//		return errors.New("OPENAI_API_KEY environment variable not set")
//	}
//	c := openai.NewClient(token)
//	ctx := context.Background()
//
//	const approxTokensPerWord = 4
//	const wordCount = 100
//	requiredTokens := approxTokensPerWord * wordCount
//
//	var historyMessages []openai.ChatCompletionMessage
//
//	message := openai.ChatCompletionMessage{
//		Role:    openai.ChatMessageRoleUser,
//		Content: usermessage,
//	}
//
//	historyMessages = append(historyMessages, message)
//
//	req := openai.ChatCompletionRequest{
//		Model:     openai.GPT3Dot5Turbo,
//		MaxTokens: requiredTokens, // Updated to accommodate a 1000-word story
//		Messages:  historyMessages,
//		Stream:    true,
//	}
//	stream, err := c.CreateChatCompletionStream(ctx, req)
//	if err != nil {
//		return fmt.Errorf("ChatCompletionStream error: %v\n", err)
//	}
//	defer stream.Close()
//
//	fmt.Printf("Stream response: ")
//	for {
//		response, err := stream.Recv()
//		if errors.Is(err, io.EOF) {
//			fmt.Println()
//			return nil
//		}
//
//		if err != nil {
//			return fmt.Errorf("\nStream error: %v\n", err)
//		}
//
//		for _, message := range response.Choices {
//			fmt.Printf(message.Delta.Content)
//		}
//	}
//
//}
